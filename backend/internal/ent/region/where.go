// Code generated by ent, DO NOT EDIT.

package region

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldDeletedAt, v))
}

// IsEnable applies equality check predicate on the "is_enable" field. It's identical to IsEnableEQ.
func IsEnable(v bool) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldIsEnable, v))
}

// NameID applies equality check predicate on the "name_id" field. It's identical to NameIDEQ.
func NameID(v string) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldNameID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldName, v))
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldPid, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldDeletedAt, v))
}

// IsEnableEQ applies the EQ predicate on the "is_enable" field.
func IsEnableEQ(v bool) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldIsEnable, v))
}

// IsEnableNEQ applies the NEQ predicate on the "is_enable" field.
func IsEnableNEQ(v bool) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldIsEnable, v))
}

// NameIDEQ applies the EQ predicate on the "name_id" field.
func NameIDEQ(v string) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldNameID, v))
}

// NameIDNEQ applies the NEQ predicate on the "name_id" field.
func NameIDNEQ(v string) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldNameID, v))
}

// NameIDIn applies the In predicate on the "name_id" field.
func NameIDIn(vs ...string) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldNameID, vs...))
}

// NameIDNotIn applies the NotIn predicate on the "name_id" field.
func NameIDNotIn(vs ...string) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldNameID, vs...))
}

// NameIDGT applies the GT predicate on the "name_id" field.
func NameIDGT(v string) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldNameID, v))
}

// NameIDGTE applies the GTE predicate on the "name_id" field.
func NameIDGTE(v string) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldNameID, v))
}

// NameIDLT applies the LT predicate on the "name_id" field.
func NameIDLT(v string) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldNameID, v))
}

// NameIDLTE applies the LTE predicate on the "name_id" field.
func NameIDLTE(v string) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldNameID, v))
}

// NameIDContains applies the Contains predicate on the "name_id" field.
func NameIDContains(v string) predicate.Region {
	return predicate.Region(sql.FieldContains(FieldNameID, v))
}

// NameIDHasPrefix applies the HasPrefix predicate on the "name_id" field.
func NameIDHasPrefix(v string) predicate.Region {
	return predicate.Region(sql.FieldHasPrefix(FieldNameID, v))
}

// NameIDHasSuffix applies the HasSuffix predicate on the "name_id" field.
func NameIDHasSuffix(v string) predicate.Region {
	return predicate.Region(sql.FieldHasSuffix(FieldNameID, v))
}

// NameIDEqualFold applies the EqualFold predicate on the "name_id" field.
func NameIDEqualFold(v string) predicate.Region {
	return predicate.Region(sql.FieldEqualFold(FieldNameID, v))
}

// NameIDContainsFold applies the ContainsFold predicate on the "name_id" field.
func NameIDContainsFold(v string) predicate.Region {
	return predicate.Region(sql.FieldContainsFold(FieldNameID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Region {
	return predicate.Region(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Region {
	return predicate.Region(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Region {
	return predicate.Region(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Region {
	return predicate.Region(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Region {
	return predicate.Region(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldType, vs...))
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldEQ(FieldPid, v))
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldNEQ(FieldPid, v))
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...xid.ID) predicate.Region {
	return predicate.Region(sql.FieldIn(FieldPid, vs...))
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...xid.ID) predicate.Region {
	return predicate.Region(sql.FieldNotIn(FieldPid, vs...))
}

// PidGT applies the GT predicate on the "pid" field.
func PidGT(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldGT(FieldPid, v))
}

// PidGTE applies the GTE predicate on the "pid" field.
func PidGTE(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldGTE(FieldPid, v))
}

// PidLT applies the LT predicate on the "pid" field.
func PidLT(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldLT(FieldPid, v))
}

// PidLTE applies the LTE predicate on the "pid" field.
func PidLTE(v xid.ID) predicate.Region {
	return predicate.Region(sql.FieldLTE(FieldPid, v))
}

// PidContains applies the Contains predicate on the "pid" field.
func PidContains(v xid.ID) predicate.Region {
	vc := string(v)
	return predicate.Region(sql.FieldContains(FieldPid, vc))
}

// PidHasPrefix applies the HasPrefix predicate on the "pid" field.
func PidHasPrefix(v xid.ID) predicate.Region {
	vc := string(v)
	return predicate.Region(sql.FieldHasPrefix(FieldPid, vc))
}

// PidHasSuffix applies the HasSuffix predicate on the "pid" field.
func PidHasSuffix(v xid.ID) predicate.Region {
	vc := string(v)
	return predicate.Region(sql.FieldHasSuffix(FieldPid, vc))
}

// PidIsNil applies the IsNil predicate on the "pid" field.
func PidIsNil() predicate.Region {
	return predicate.Region(sql.FieldIsNull(FieldPid))
}

// PidNotNil applies the NotNil predicate on the "pid" field.
func PidNotNil() predicate.Region {
	return predicate.Region(sql.FieldNotNull(FieldPid))
}

// PidEqualFold applies the EqualFold predicate on the "pid" field.
func PidEqualFold(v xid.ID) predicate.Region {
	vc := string(v)
	return predicate.Region(sql.FieldEqualFold(FieldPid, vc))
}

// PidContainsFold applies the ContainsFold predicate on the "pid" field.
func PidContainsFold(v xid.ID) predicate.Region {
	vc := string(v)
	return predicate.Region(sql.FieldContainsFold(FieldPid, vc))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Region {
	return predicate.Region(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Region
		step.Edge.Schema = schemaConfig.Region
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Region) predicate.Region {
	return predicate.Region(func(s *sql.Selector) {
		step := newParentStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Region
		step.Edge.Schema = schemaConfig.Region
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Region {
	return predicate.Region(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Region
		step.Edge.Schema = schemaConfig.Region
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Region) predicate.Region {
	return predicate.Region(func(s *sql.Selector) {
		step := newChildrenStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Region
		step.Edge.Schema = schemaConfig.Region
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Region) predicate.Region {
	return predicate.Region(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Region) predicate.Region {
	return predicate.Region(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Region) predicate.Region {
	return predicate.Region(sql.NotPredicates(p))
}
