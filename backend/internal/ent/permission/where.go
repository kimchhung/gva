// Code generated by ent, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pxid.ID) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldGroup, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldScope, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrder, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldUpdatedAt, v))
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldGroup, v))
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldGroup, v))
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldGroup, v))
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldGroup, v))
}

// GroupEqualFold applies the EqualFold predicate on the "group" field.
func GroupEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldGroup, v))
}

// GroupContainsFold applies the ContainsFold predicate on the "group" field.
func GroupContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldGroup, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldName, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldScope, vs...))
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldScope, v))
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldScope, v))
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldScope, v))
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldScope, v))
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldScope, v))
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldScope, v))
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldScope, v))
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldScope, v))
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldScope, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldType, vs...))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldType))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldOrder, v))
}

// OrderIsNil applies the IsNil predicate on the "order" field.
func OrderIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldOrder))
}

// OrderNotNil applies the NotNil predicate on the "order" field.
func OrderNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldOrder))
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.RolePermissions
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newRolesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.RolePermissions
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.NotPredicates(p))
}
