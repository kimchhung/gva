// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/internal/ent/admin"
	"github.com/kimchhung/gva/internal/ent/role"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AdminCreate) SetCreatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AdminCreate) SetUpdatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetUsername sets the "username" field.
func (ac *AdminCreate) SetUsername(s string) *AdminCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AdminCreate) SetPassword(s string) *AdminCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetWhitelistIps sets the "whitelist_ips" field.
func (ac *AdminCreate) SetWhitelistIps(s []string) *AdminCreate {
	ac.mutation.SetWhitelistIps(s)
	return ac
}

// SetIsActive sets the "is_active" field.
func (ac *AdminCreate) SetIsActive(b bool) *AdminCreate {
	ac.mutation.SetIsActive(b)
	return ac
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ac *AdminCreate) SetNillableIsActive(b *bool) *AdminCreate {
	if b != nil {
		ac.SetIsActive(*b)
	}
	return ac
}

// SetDisplayName sets the "display_name" field.
func (ac *AdminCreate) SetDisplayName(s string) *AdminCreate {
	ac.mutation.SetDisplayName(s)
	return ac
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ac *AdminCreate) SetNillableDisplayName(s *string) *AdminCreate {
	if s != nil {
		ac.SetDisplayName(*s)
	}
	return ac
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (ac *AdminCreate) AddRoleIDs(ids ...int) *AdminCreate {
	ac.mutation.AddRoleIDs(ids...)
	return ac
}

// AddRoles adds the "roles" edges to the Role entity.
func (ac *AdminCreate) AddRoles(r ...*Role) *AdminCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddRoleIDs(ids...)
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := admin.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := admin.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.IsActive(); !ok {
		v := admin.DefaultIsActive
		ac.mutation.SetIsActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Admin.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Admin.updated_at"`)}
	}
	if _, ok := ac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Admin.username"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Admin.password"`)}
	}
	if _, ok := ac.mutation.WhitelistIps(); !ok {
		return &ValidationError{Name: "whitelist_ips", err: errors.New(`ent: missing required field "Admin.whitelist_ips"`)}
	}
	if _, ok := ac.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Admin.is_active"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(admin.Table, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.WhitelistIps(); ok {
		_spec.SetField(admin.FieldWhitelistIps, field.TypeJSON, value)
		_node.WhitelistIps = value
	}
	if value, ok := ac.mutation.IsActive(); ok {
		_spec.SetField(admin.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := ac.mutation.DisplayName(); ok {
		_spec.SetField(admin.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if nodes := ac.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	err      error
	builders []*AdminCreate
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
