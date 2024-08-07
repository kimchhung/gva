// Code generated by ent, DO NOT EDIT.

package menu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent/predicate"

	"github.com/gva/internal/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsEnable applies equality check predicate on the "is_enable" field. It's identical to IsEnableEQ.
func IsEnable(v bool) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldIsEnable, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldDeletedAt, v))
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldPid, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldPath, v))
}

// Component applies equality check predicate on the "component" field. It's identical to ComponentEQ.
func Component(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldComponent, v))
}

// Redirect applies equality check predicate on the "redirect" field. It's identical to RedirectEQ.
func Redirect(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldRedirect, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldName, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldOrder, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsEnableEQ applies the EQ predicate on the "is_enable" field.
func IsEnableEQ(v bool) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldIsEnable, v))
}

// IsEnableNEQ applies the NEQ predicate on the "is_enable" field.
func IsEnableNEQ(v bool) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldIsEnable, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldDeletedAt, v))
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldPid, v))
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldPid, v))
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldPid, vs...))
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldPid, vs...))
}

// PidGT applies the GT predicate on the "pid" field.
func PidGT(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldPid, v))
}

// PidGTE applies the GTE predicate on the "pid" field.
func PidGTE(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldPid, v))
}

// PidLT applies the LT predicate on the "pid" field.
func PidLT(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldPid, v))
}

// PidLTE applies the LTE predicate on the "pid" field.
func PidLTE(v pxid.ID) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldPid, v))
}

// PidContains applies the Contains predicate on the "pid" field.
func PidContains(v pxid.ID) predicate.Menu {
	vc := string(v)
	return predicate.Menu(sql.FieldContains(FieldPid, vc))
}

// PidHasPrefix applies the HasPrefix predicate on the "pid" field.
func PidHasPrefix(v pxid.ID) predicate.Menu {
	vc := string(v)
	return predicate.Menu(sql.FieldHasPrefix(FieldPid, vc))
}

// PidHasSuffix applies the HasSuffix predicate on the "pid" field.
func PidHasSuffix(v pxid.ID) predicate.Menu {
	vc := string(v)
	return predicate.Menu(sql.FieldHasSuffix(FieldPid, vc))
}

// PidIsNil applies the IsNil predicate on the "pid" field.
func PidIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldPid))
}

// PidNotNil applies the NotNil predicate on the "pid" field.
func PidNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldPid))
}

// PidEqualFold applies the EqualFold predicate on the "pid" field.
func PidEqualFold(v pxid.ID) predicate.Menu {
	vc := string(v)
	return predicate.Menu(sql.FieldEqualFold(FieldPid, vc))
}

// PidContainsFold applies the ContainsFold predicate on the "pid" field.
func PidContainsFold(v pxid.ID) predicate.Menu {
	vc := string(v)
	return predicate.Menu(sql.FieldContainsFold(FieldPid, vc))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldPath, v))
}

// ComponentEQ applies the EQ predicate on the "component" field.
func ComponentEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldComponent, v))
}

// ComponentNEQ applies the NEQ predicate on the "component" field.
func ComponentNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldComponent, v))
}

// ComponentIn applies the In predicate on the "component" field.
func ComponentIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldComponent, vs...))
}

// ComponentNotIn applies the NotIn predicate on the "component" field.
func ComponentNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldComponent, vs...))
}

// ComponentGT applies the GT predicate on the "component" field.
func ComponentGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldComponent, v))
}

// ComponentGTE applies the GTE predicate on the "component" field.
func ComponentGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldComponent, v))
}

// ComponentLT applies the LT predicate on the "component" field.
func ComponentLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldComponent, v))
}

// ComponentLTE applies the LTE predicate on the "component" field.
func ComponentLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldComponent, v))
}

// ComponentContains applies the Contains predicate on the "component" field.
func ComponentContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldComponent, v))
}

// ComponentHasPrefix applies the HasPrefix predicate on the "component" field.
func ComponentHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldComponent, v))
}

// ComponentHasSuffix applies the HasSuffix predicate on the "component" field.
func ComponentHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldComponent, v))
}

// ComponentEqualFold applies the EqualFold predicate on the "component" field.
func ComponentEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldComponent, v))
}

// ComponentContainsFold applies the ContainsFold predicate on the "component" field.
func ComponentContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldComponent, v))
}

// RedirectEQ applies the EQ predicate on the "redirect" field.
func RedirectEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldRedirect, v))
}

// RedirectNEQ applies the NEQ predicate on the "redirect" field.
func RedirectNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldRedirect, v))
}

// RedirectIn applies the In predicate on the "redirect" field.
func RedirectIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldRedirect, vs...))
}

// RedirectNotIn applies the NotIn predicate on the "redirect" field.
func RedirectNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldRedirect, vs...))
}

// RedirectGT applies the GT predicate on the "redirect" field.
func RedirectGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldRedirect, v))
}

// RedirectGTE applies the GTE predicate on the "redirect" field.
func RedirectGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldRedirect, v))
}

// RedirectLT applies the LT predicate on the "redirect" field.
func RedirectLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldRedirect, v))
}

// RedirectLTE applies the LTE predicate on the "redirect" field.
func RedirectLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldRedirect, v))
}

// RedirectContains applies the Contains predicate on the "redirect" field.
func RedirectContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldRedirect, v))
}

// RedirectHasPrefix applies the HasPrefix predicate on the "redirect" field.
func RedirectHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldRedirect, v))
}

// RedirectHasSuffix applies the HasSuffix predicate on the "redirect" field.
func RedirectHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldRedirect, v))
}

// RedirectIsNil applies the IsNil predicate on the "redirect" field.
func RedirectIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldRedirect))
}

// RedirectNotNil applies the NotNil predicate on the "redirect" field.
func RedirectNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldRedirect))
}

// RedirectEqualFold applies the EqualFold predicate on the "redirect" field.
func RedirectEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldRedirect, v))
}

// RedirectContainsFold applies the ContainsFold predicate on the "redirect" field.
func RedirectContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldRedirect, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldName, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldOrder, v))
}

// OrderIsNil applies the IsNil predicate on the "order" field.
func OrderIsNil() predicate.Menu {
	return predicate.Menu(sql.FieldIsNull(FieldOrder))
}

// OrderNotNil applies the NotNil predicate on the "order" field.
func OrderNotNil() predicate.Menu {
	return predicate.Menu(sql.FieldNotNull(FieldOrder))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldType, vs...))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Menu
		step.Edge.Schema = schemaConfig.Menu
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := newParentStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Menu
		step.Edge.Schema = schemaConfig.Menu
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Menu
		step.Edge.Schema = schemaConfig.Menu
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Menu) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := newChildrenStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Menu
		step.Edge.Schema = schemaConfig.Menu
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.RoleRoutes
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Menu {
	return predicate.Menu(func(s *sql.Selector) {
		step := newRolesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.RoleRoutes
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.NotPredicates(p))
}
