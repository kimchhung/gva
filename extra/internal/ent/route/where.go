// Code generated by ent, DO NOT EDIT.

package route

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kimchhung/gva/extra/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsEnable applies equality check predicate on the "is_enable" field. It's identical to IsEnableEQ.
func IsEnable(v bool) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldIsEnable, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldDeletedAt, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldParentID, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldPath, v))
}

// Component applies equality check predicate on the "component" field. It's identical to ComponentEQ.
func Component(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldComponent, v))
}

// Redirect applies equality check predicate on the "redirect" field. It's identical to RedirectEQ.
func Redirect(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldRedirect, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldType, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldTitle, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsEnableEQ applies the EQ predicate on the "is_enable" field.
func IsEnableEQ(v bool) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldIsEnable, v))
}

// IsEnableNEQ applies the NEQ predicate on the "is_enable" field.
func IsEnableNEQ(v bool) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldIsEnable, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldDeletedAt, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Route {
	return predicate.Route(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Route {
	return predicate.Route(sql.FieldNotNull(FieldParentID))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Route {
	return predicate.Route(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Route {
	return predicate.Route(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Route {
	return predicate.Route(sql.FieldContainsFold(FieldPath, v))
}

// ComponentEQ applies the EQ predicate on the "component" field.
func ComponentEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldComponent, v))
}

// ComponentNEQ applies the NEQ predicate on the "component" field.
func ComponentNEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldComponent, v))
}

// ComponentIn applies the In predicate on the "component" field.
func ComponentIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldComponent, vs...))
}

// ComponentNotIn applies the NotIn predicate on the "component" field.
func ComponentNotIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldComponent, vs...))
}

// ComponentGT applies the GT predicate on the "component" field.
func ComponentGT(v string) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldComponent, v))
}

// ComponentGTE applies the GTE predicate on the "component" field.
func ComponentGTE(v string) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldComponent, v))
}

// ComponentLT applies the LT predicate on the "component" field.
func ComponentLT(v string) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldComponent, v))
}

// ComponentLTE applies the LTE predicate on the "component" field.
func ComponentLTE(v string) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldComponent, v))
}

// ComponentContains applies the Contains predicate on the "component" field.
func ComponentContains(v string) predicate.Route {
	return predicate.Route(sql.FieldContains(FieldComponent, v))
}

// ComponentHasPrefix applies the HasPrefix predicate on the "component" field.
func ComponentHasPrefix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasPrefix(FieldComponent, v))
}

// ComponentHasSuffix applies the HasSuffix predicate on the "component" field.
func ComponentHasSuffix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasSuffix(FieldComponent, v))
}

// ComponentEqualFold applies the EqualFold predicate on the "component" field.
func ComponentEqualFold(v string) predicate.Route {
	return predicate.Route(sql.FieldEqualFold(FieldComponent, v))
}

// ComponentContainsFold applies the ContainsFold predicate on the "component" field.
func ComponentContainsFold(v string) predicate.Route {
	return predicate.Route(sql.FieldContainsFold(FieldComponent, v))
}

// RedirectEQ applies the EQ predicate on the "redirect" field.
func RedirectEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldRedirect, v))
}

// RedirectNEQ applies the NEQ predicate on the "redirect" field.
func RedirectNEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldRedirect, v))
}

// RedirectIn applies the In predicate on the "redirect" field.
func RedirectIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldRedirect, vs...))
}

// RedirectNotIn applies the NotIn predicate on the "redirect" field.
func RedirectNotIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldRedirect, vs...))
}

// RedirectGT applies the GT predicate on the "redirect" field.
func RedirectGT(v string) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldRedirect, v))
}

// RedirectGTE applies the GTE predicate on the "redirect" field.
func RedirectGTE(v string) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldRedirect, v))
}

// RedirectLT applies the LT predicate on the "redirect" field.
func RedirectLT(v string) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldRedirect, v))
}

// RedirectLTE applies the LTE predicate on the "redirect" field.
func RedirectLTE(v string) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldRedirect, v))
}

// RedirectContains applies the Contains predicate on the "redirect" field.
func RedirectContains(v string) predicate.Route {
	return predicate.Route(sql.FieldContains(FieldRedirect, v))
}

// RedirectHasPrefix applies the HasPrefix predicate on the "redirect" field.
func RedirectHasPrefix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasPrefix(FieldRedirect, v))
}

// RedirectHasSuffix applies the HasSuffix predicate on the "redirect" field.
func RedirectHasSuffix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasSuffix(FieldRedirect, v))
}

// RedirectEqualFold applies the EqualFold predicate on the "redirect" field.
func RedirectEqualFold(v string) predicate.Route {
	return predicate.Route(sql.FieldEqualFold(FieldRedirect, v))
}

// RedirectContainsFold applies the ContainsFold predicate on the "redirect" field.
func RedirectContainsFold(v string) predicate.Route {
	return predicate.Route(sql.FieldContainsFold(FieldRedirect, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Route {
	return predicate.Route(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Route {
	return predicate.Route(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Route {
	return predicate.Route(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldType, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Route {
	return predicate.Route(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Route {
	return predicate.Route(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Route {
	return predicate.Route(sql.FieldContainsFold(FieldTitle, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		step := newRolesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Route) predicate.Route {
	return predicate.Route(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Route) predicate.Route {
	return predicate.Route(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Route) predicate.Route {
	return predicate.Route(sql.NotPredicates(p))
}
