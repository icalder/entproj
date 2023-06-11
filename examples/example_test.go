package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/icalder/entproj/util"
	_ "github.com/lib/pq"
)

func ExampleRegistry() {
	ctx := context.Background()
	client, tx := util.OpenSQLiteAndStartTransaction(ctx)
	defer client.Close()
	defer tx.Rollback()

	// Run the automatic migration tool to create all schema resources.
	if err := tx.Client().Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	registry1, err := tx.Registry.Create().
		SetID(uuid.MustParse("1e325c56-c609-4c17-b68b-474eb9097684")).
		SetName("my-registry").
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating a registry: %v", err)
	}

	registry2, err := tx.Registry.Create().
		SetID(uuid.MustParse("10d5308c-6a54-463c-adb6-7c9cef932cea")).
		SetName("test-registry").
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating a registry: %v", err)
	}

	repo1, err := tx.Repository.Create().
		SetName("ubuntu").
		SetRegistry(registry1).
		Save(ctx)

	fmt.Println(registry1)
	fmt.Println(registry2)
	fmt.Println(repo1.Name)
	// Output:
	// Registry(id=1e325c56-c609-4c17-b68b-474eb9097684, name=my-registry)
	// Registry(id=10d5308c-6a54-463c-adb6-7c9cef932cea, name=test-registry)
	// ubuntu
}
