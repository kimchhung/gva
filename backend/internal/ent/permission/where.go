// Code generated by ent, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kimchhung/gva/backend/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
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

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldKey, v))
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

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldKey, v))
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

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newRolesStep()
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
