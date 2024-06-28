// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/role"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *PermissionCreate) SetCreatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PermissionCreate) SetUpdatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetGroup sets the "group" field.
func (pc *PermissionCreate) SetGroup(s string) *PermissionCreate {
	pc.mutation.SetGroup(s)
	return pc
}

// SetName sets the "name" field.
func (pc *PermissionCreate) SetName(s string) *PermissionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetKey sets the "key" field.
func (pc *PermissionCreate) SetKey(s string) *PermissionCreate {
	pc.mutation.SetKey(s)
	return pc
}

// SetOrder sets the "order" field.
func (pc *PermissionCreate) SetOrder(i int) *PermissionCreate {
	pc.mutation.SetOrder(i)
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(pu pulid.ID) *PermissionCreate {
	pc.mutation.SetID(pu)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableID(pu *pulid.ID) *PermissionCreate {
	if pu != nil {
		pc.SetID(*pu)
	}
	return pc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pc *PermissionCreate) AddRoleIDs(ids ...pulid.ID) *PermissionCreate {
	pc.mutation.AddRoleIDs(ids...)
	return pc
}

// AddRoles adds the "roles" edges to the Role entity.
func (pc *PermissionCreate) AddRoles(r ...*Role) *PermissionCreate {
	ids := make([]pulid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermissionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermissionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := permission.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := permission.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := permission.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Permission.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Permission.updated_at"`)}
	}
	if _, ok := pc.mutation.Group(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required field "Permission.group"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Permission.name"`)}
	}
	if _, ok := pc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Permission.key"`)}
	}
	if _, ok := pc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Permission.order"`)}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(permission.Table, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeString))
	)
	_spec.Schema = pc.schemaConfig.Permission
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(permission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Group(); ok {
		_spec.SetField(permission.FieldGroup, field.TypeString, value)
		_node.Group = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Key(); ok {
		_spec.SetField(permission.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := pc.mutation.Order(); ok {
		_spec.SetField(permission.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := pc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.RolePermissions
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Permission.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PermissionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pc *PermissionCreate) OnConflict(opts ...sql.ConflictOption) *PermissionUpsertOne {
	pc.conflict = opts
	return &PermissionUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PermissionCreate) OnConflictColumns(columns ...string) *PermissionUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PermissionUpsertOne{
		create: pc,
	}
}

type (
	// PermissionUpsertOne is the builder for "upsert"-ing
	//  one Permission node.
	PermissionUpsertOne struct {
		create *PermissionCreate
	}

	// PermissionUpsert is the "OnConflict" setter.
	PermissionUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PermissionUpsert) SetCreatedAt(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateCreatedAt() *PermissionUpsert {
	u.SetExcluded(permission.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PermissionUpsert) SetUpdatedAt(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateUpdatedAt() *PermissionUpsert {
	u.SetExcluded(permission.FieldUpdatedAt)
	return u
}

// SetGroup sets the "group" field.
func (u *PermissionUpsert) SetGroup(v string) *PermissionUpsert {
	u.Set(permission.FieldGroup, v)
	return u
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateGroup() *PermissionUpsert {
	u.SetExcluded(permission.FieldGroup)
	return u
}

// SetName sets the "name" field.
func (u *PermissionUpsert) SetName(v string) *PermissionUpsert {
	u.Set(permission.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateName() *PermissionUpsert {
	u.SetExcluded(permission.FieldName)
	return u
}

// SetKey sets the "key" field.
func (u *PermissionUpsert) SetKey(v string) *PermissionUpsert {
	u.Set(permission.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateKey() *PermissionUpsert {
	u.SetExcluded(permission.FieldKey)
	return u
}

// SetOrder sets the "order" field.
func (u *PermissionUpsert) SetOrder(v int) *PermissionUpsert {
	u.Set(permission.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateOrder() *PermissionUpsert {
	u.SetExcluded(permission.FieldOrder)
	return u
}

// AddOrder adds v to the "order" field.
func (u *PermissionUpsert) AddOrder(v int) *PermissionUpsert {
	u.Add(permission.FieldOrder, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(permission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PermissionUpsertOne) UpdateNewValues() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(permission.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PermissionUpsertOne) Ignore() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PermissionUpsertOne) DoNothing() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PermissionCreate.OnConflict
// documentation for more info.
func (u *PermissionUpsertOne) Update(set func(*PermissionUpsert)) *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PermissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PermissionUpsertOne) SetCreatedAt(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateCreatedAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PermissionUpsertOne) SetUpdatedAt(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateUpdatedAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetGroup sets the "group" field.
func (u *PermissionUpsertOne) SetGroup(v string) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetGroup(v)
	})
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateGroup() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateGroup()
	})
}

// SetName sets the "name" field.
func (u *PermissionUpsertOne) SetName(v string) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateName() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateName()
	})
}

// SetKey sets the "key" field.
func (u *PermissionUpsertOne) SetKey(v string) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateKey() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateKey()
	})
}

// SetOrder sets the "order" field.
func (u *PermissionUpsertOne) SetOrder(v int) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *PermissionUpsertOne) AddOrder(v int) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateOrder() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateOrder()
	})
}

// Exec executes the query.
func (u *PermissionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PermissionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PermissionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PermissionUpsertOne) ID(ctx context.Context) (id pulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PermissionUpsertOne.ID is not supported by MySQL driver. Use PermissionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PermissionUpsertOne) IDX(ctx context.Context) pulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	err      error
	builders []*PermissionCreate
	conflict []sql.ConflictOption
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermissionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Permission.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PermissionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pcb *PermissionCreateBulk) OnConflict(opts ...sql.ConflictOption) *PermissionUpsertBulk {
	pcb.conflict = opts
	return &PermissionUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PermissionCreateBulk) OnConflictColumns(columns ...string) *PermissionUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PermissionUpsertBulk{
		create: pcb,
	}
}

// PermissionUpsertBulk is the builder for "upsert"-ing
// a bulk of Permission nodes.
type PermissionUpsertBulk struct {
	create *PermissionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(permission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PermissionUpsertBulk) UpdateNewValues() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(permission.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PermissionUpsertBulk) Ignore() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PermissionUpsertBulk) DoNothing() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PermissionCreateBulk.OnConflict
// documentation for more info.
func (u *PermissionUpsertBulk) Update(set func(*PermissionUpsert)) *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PermissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PermissionUpsertBulk) SetCreatedAt(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateCreatedAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PermissionUpsertBulk) SetUpdatedAt(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateUpdatedAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetGroup sets the "group" field.
func (u *PermissionUpsertBulk) SetGroup(v string) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetGroup(v)
	})
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateGroup() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateGroup()
	})
}

// SetName sets the "name" field.
func (u *PermissionUpsertBulk) SetName(v string) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateName() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateName()
	})
}

// SetKey sets the "key" field.
func (u *PermissionUpsertBulk) SetKey(v string) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateKey() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateKey()
	})
}

// SetOrder sets the "order" field.
func (u *PermissionUpsertBulk) SetOrder(v int) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *PermissionUpsertBulk) AddOrder(v int) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateOrder() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateOrder()
	})
}

// Exec executes the query.
func (u *PermissionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PermissionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PermissionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PermissionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
