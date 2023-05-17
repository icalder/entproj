// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/icalder/enttest/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/icalder/enttest/ent/registry"
	"github.com/icalder/enttest/ent/repository"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Registry is the client for interacting with the Registry builders.
	Registry *RegistryClient
	// Repository is the client for interacting with the Repository builders.
	Repository *RepositoryClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Registry = NewRegistryClient(c.config)
	c.Repository = NewRepositoryClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Registry:   NewRegistryClient(cfg),
		Repository: NewRepositoryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Registry:   NewRegistryClient(cfg),
		Repository: NewRepositoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Registry.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Registry.Use(hooks...)
	c.Repository.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Registry.Intercept(interceptors...)
	c.Repository.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *RegistryMutation:
		return c.Registry.mutate(ctx, m)
	case *RepositoryMutation:
		return c.Repository.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// RegistryClient is a client for the Registry schema.
type RegistryClient struct {
	config
}

// NewRegistryClient returns a client for the Registry from the given config.
func NewRegistryClient(c config) *RegistryClient {
	return &RegistryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `registry.Hooks(f(g(h())))`.
func (c *RegistryClient) Use(hooks ...Hook) {
	c.hooks.Registry = append(c.hooks.Registry, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `registry.Intercept(f(g(h())))`.
func (c *RegistryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Registry = append(c.inters.Registry, interceptors...)
}

// Create returns a builder for creating a Registry entity.
func (c *RegistryClient) Create() *RegistryCreate {
	mutation := newRegistryMutation(c.config, OpCreate)
	return &RegistryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Registry entities.
func (c *RegistryClient) CreateBulk(builders ...*RegistryCreate) *RegistryCreateBulk {
	return &RegistryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Registry.
func (c *RegistryClient) Update() *RegistryUpdate {
	mutation := newRegistryMutation(c.config, OpUpdate)
	return &RegistryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RegistryClient) UpdateOne(r *Registry) *RegistryUpdateOne {
	mutation := newRegistryMutation(c.config, OpUpdateOne, withRegistry(r))
	return &RegistryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RegistryClient) UpdateOneID(id uuid.UUID) *RegistryUpdateOne {
	mutation := newRegistryMutation(c.config, OpUpdateOne, withRegistryID(id))
	return &RegistryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Registry.
func (c *RegistryClient) Delete() *RegistryDelete {
	mutation := newRegistryMutation(c.config, OpDelete)
	return &RegistryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RegistryClient) DeleteOne(r *Registry) *RegistryDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RegistryClient) DeleteOneID(id uuid.UUID) *RegistryDeleteOne {
	builder := c.Delete().Where(registry.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RegistryDeleteOne{builder}
}

// Query returns a query builder for Registry.
func (c *RegistryClient) Query() *RegistryQuery {
	return &RegistryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRegistry},
		inters: c.Interceptors(),
	}
}

// Get returns a Registry entity by its id.
func (c *RegistryClient) Get(ctx context.Context, id uuid.UUID) (*Registry, error) {
	return c.Query().Where(registry.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RegistryClient) GetX(ctx context.Context, id uuid.UUID) *Registry {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRepositories queries the repositories edge of a Registry.
func (c *RegistryClient) QueryRepositories(r *Registry) *RepositoryQuery {
	query := (&RepositoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(registry.Table, registry.FieldID, id),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, registry.RepositoriesTable, registry.RepositoriesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RegistryClient) Hooks() []Hook {
	return c.hooks.Registry
}

// Interceptors returns the client interceptors.
func (c *RegistryClient) Interceptors() []Interceptor {
	return c.inters.Registry
}

func (c *RegistryClient) mutate(ctx context.Context, m *RegistryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RegistryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RegistryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RegistryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RegistryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Registry mutation op: %q", m.Op())
	}
}

// RepositoryClient is a client for the Repository schema.
type RepositoryClient struct {
	config
}

// NewRepositoryClient returns a client for the Repository from the given config.
func NewRepositoryClient(c config) *RepositoryClient {
	return &RepositoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `repository.Hooks(f(g(h())))`.
func (c *RepositoryClient) Use(hooks ...Hook) {
	c.hooks.Repository = append(c.hooks.Repository, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `repository.Intercept(f(g(h())))`.
func (c *RepositoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Repository = append(c.inters.Repository, interceptors...)
}

// Create returns a builder for creating a Repository entity.
func (c *RepositoryClient) Create() *RepositoryCreate {
	mutation := newRepositoryMutation(c.config, OpCreate)
	return &RepositoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Repository entities.
func (c *RepositoryClient) CreateBulk(builders ...*RepositoryCreate) *RepositoryCreateBulk {
	return &RepositoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Repository.
func (c *RepositoryClient) Update() *RepositoryUpdate {
	mutation := newRepositoryMutation(c.config, OpUpdate)
	return &RepositoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RepositoryClient) UpdateOne(r *Repository) *RepositoryUpdateOne {
	mutation := newRepositoryMutation(c.config, OpUpdateOne, withRepository(r))
	return &RepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RepositoryClient) UpdateOneID(id int) *RepositoryUpdateOne {
	mutation := newRepositoryMutation(c.config, OpUpdateOne, withRepositoryID(id))
	return &RepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Repository.
func (c *RepositoryClient) Delete() *RepositoryDelete {
	mutation := newRepositoryMutation(c.config, OpDelete)
	return &RepositoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RepositoryClient) DeleteOne(r *Repository) *RepositoryDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RepositoryClient) DeleteOneID(id int) *RepositoryDeleteOne {
	builder := c.Delete().Where(repository.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RepositoryDeleteOne{builder}
}

// Query returns a query builder for Repository.
func (c *RepositoryClient) Query() *RepositoryQuery {
	return &RepositoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRepository},
		inters: c.Interceptors(),
	}
}

// Get returns a Repository entity by its id.
func (c *RepositoryClient) Get(ctx context.Context, id int) (*Repository, error) {
	return c.Query().Where(repository.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RepositoryClient) GetX(ctx context.Context, id int) *Repository {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRegistry queries the registry edge of a Repository.
func (c *RepositoryClient) QueryRegistry(r *Repository) *RegistryQuery {
	query := (&RegistryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(repository.Table, repository.FieldID, id),
			sqlgraph.To(registry.Table, registry.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, repository.RegistryTable, repository.RegistryColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RepositoryClient) Hooks() []Hook {
	return c.hooks.Repository
}

// Interceptors returns the client interceptors.
func (c *RepositoryClient) Interceptors() []Interceptor {
	return c.inters.Repository
}

func (c *RepositoryClient) mutate(ctx context.Context, m *RepositoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RepositoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RepositoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RepositoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Repository mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Registry, Repository []ent.Hook
	}
	inters struct {
		Registry, Repository []ent.Interceptor
	}
)
