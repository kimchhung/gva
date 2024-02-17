// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/kimchhung/gva/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kimchhung/gva/internal/ent/admin"
	"github.com/kimchhung/gva/internal/ent/permission"
	"github.com/kimchhung/gva/internal/ent/role"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Admin is the client for interacting with the Admin builders.
	Admin *AdminClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Admin = NewAdminClient(c.config)
	c.Permission = NewPermissionClient(c.config)
	c.Role = NewRoleClient(c.config)
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
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
		Admin:      NewAdminClient(cfg),
		Permission: NewPermissionClient(cfg),
		Role:       NewRoleClient(cfg),
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
		Admin:      NewAdminClient(cfg),
		Permission: NewPermissionClient(cfg),
		Role:       NewRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Admin.
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
	c.Admin.Use(hooks...)
	c.Permission.Use(hooks...)
	c.Role.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Admin.Intercept(interceptors...)
	c.Permission.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AdminMutation:
		return c.Admin.mutate(ctx, m)
	case *PermissionMutation:
		return c.Permission.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AdminClient is a client for the Admin schema.
type AdminClient struct {
	config
}

// NewAdminClient returns a client for the Admin from the given config.
func NewAdminClient(c config) *AdminClient {
	return &AdminClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `admin.Hooks(f(g(h())))`.
func (c *AdminClient) Use(hooks ...Hook) {
	c.hooks.Admin = append(c.hooks.Admin, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `admin.Intercept(f(g(h())))`.
func (c *AdminClient) Intercept(interceptors ...Interceptor) {
	c.inters.Admin = append(c.inters.Admin, interceptors...)
}

// Create returns a builder for creating a Admin entity.
func (c *AdminClient) Create() *AdminCreate {
	mutation := newAdminMutation(c.config, OpCreate)
	return &AdminCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Admin entities.
func (c *AdminClient) CreateBulk(builders ...*AdminCreate) *AdminCreateBulk {
	return &AdminCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AdminClient) MapCreateBulk(slice any, setFunc func(*AdminCreate, int)) *AdminCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AdminCreateBulk{err: fmt.Errorf("calling to AdminClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AdminCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AdminCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Admin.
func (c *AdminClient) Update() *AdminUpdate {
	mutation := newAdminMutation(c.config, OpUpdate)
	return &AdminUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AdminClient) UpdateOne(a *Admin) *AdminUpdateOne {
	mutation := newAdminMutation(c.config, OpUpdateOne, withAdmin(a))
	return &AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AdminClient) UpdateOneID(id int) *AdminUpdateOne {
	mutation := newAdminMutation(c.config, OpUpdateOne, withAdminID(id))
	return &AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Admin.
func (c *AdminClient) Delete() *AdminDelete {
	mutation := newAdminMutation(c.config, OpDelete)
	return &AdminDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AdminClient) DeleteOne(a *Admin) *AdminDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AdminClient) DeleteOneID(id int) *AdminDeleteOne {
	builder := c.Delete().Where(admin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AdminDeleteOne{builder}
}

// Query returns a query builder for Admin.
func (c *AdminClient) Query() *AdminQuery {
	return &AdminQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAdmin},
		inters: c.Interceptors(),
	}
}

// Get returns a Admin entity by its id.
func (c *AdminClient) Get(ctx context.Context, id int) (*Admin, error) {
	return c.Query().Where(admin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AdminClient) GetX(ctx context.Context, id int) *Admin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Admin.
func (c *AdminClient) QueryRoles(a *Admin) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(admin.Table, admin.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, admin.RolesTable, admin.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AdminClient) Hooks() []Hook {
	return c.hooks.Admin
}

// Interceptors returns the client interceptors.
func (c *AdminClient) Interceptors() []Interceptor {
	return c.inters.Admin
}

func (c *AdminClient) mutate(ctx context.Context, m *AdminMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AdminCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AdminUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AdminDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Admin mutation op: %q", m.Op())
	}
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `permission.Intercept(f(g(h())))`.
func (c *PermissionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Permission = append(c.inters.Permission, interceptors...)
}

// Create returns a builder for creating a Permission entity.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PermissionClient) MapCreateBulk(slice any, setFunc func(*PermissionCreate, int)) *PermissionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PermissionCreateBulk{err: fmt.Errorf("calling to PermissionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PermissionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id int) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PermissionClient) DeleteOneID(id int) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePermission},
		inters: c.Interceptors(),
	}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id int) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id int) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Permission.
func (c *PermissionClient) QueryRoles(pe *Permission) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.RolesTable, permission.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// Interceptors returns the client interceptors.
func (c *PermissionClient) Interceptors() []Interceptor {
	return c.inters.Permission
}

func (c *PermissionClient) mutate(ctx context.Context, m *PermissionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Permission mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoleClient) MapCreateBulk(slice any, setFunc func(*RoleCreate, int)) *RoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoleCreateBulk{err: fmt.Errorf("calling to RoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAdmins queries the admins edge of a Role.
func (c *RoleClient) QueryAdmins(r *Role) *AdminQuery {
	query := (&AdminClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.AdminsTable, role.AdminsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermissions queries the permissions edge of a Role.
func (c *RoleClient) QueryPermissions(r *Role) *PermissionQuery {
	query := (&PermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.PermissionsTable, role.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Admin, Permission, Role []ent.Hook
	}
	inters struct {
		Admin, Permission, Role []ent.Interceptor
	}
)
