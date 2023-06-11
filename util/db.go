package util

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/icalder/entproj/ent"
	"github.com/icalder/entproj/ent/migrate"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/net/context"
	"modernc.org/sqlite"
	_ "modernc.org/sqlite"
)

func init() {
	// ent uses 'sqlite3' github.com/mattn/go-sqlite3
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	// Work around for error: no such table: sqlite_sequence
	if _, err := c.Exec("CREATE TABLE dummy (id INTEGER PRIMARY KEY AUTOINCREMENT)", nil); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to work around sqlite_sequence error: %w", err)
	}
	return conn, nil
}

// For testing
func OpenSQLiteAndStartTransaction(ctx context.Context) (*ent.Client, *ent.Tx) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatalf("starting a transaction: %v", err)
	}

	// Run the automatic migration tool to create all schema resources.
	if err := tx.Client().Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, tx
}

// Open new connection
// Example usage:
// "postgres://user:pw@localhost/dbname"
// dbConnectionString, err := getDBConnectionString()
//
//	if err != nil {
//		return err
//	}
//
// entClient, err := util.OpenDB(dbConnectionString)
//
//	if err != nil {
//		log.Err(err).Msg("failed opening connection to PostgreSQL database")
//		return err
//	}
//
// defer entClient.Close()
func OpenDB(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), err
}

func MigrateSchema(client *ent.Client) error {
	// TODO move to versioned migrations
	// See https://entgo.io/docs/versioned/intro
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return client. /*.Debug()*/ Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
