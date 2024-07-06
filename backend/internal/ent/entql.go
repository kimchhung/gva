// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/predicate"
	"github.com/gva/internal/ent/region"
	"github.com/gva/internal/ent/role"
	"github.com/gva/internal/ent/route"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 6)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: admin.FieldID,
			},
		},
		Type: "Admin",
		Fields: map[string]*sqlgraph.FieldSpec{
			admin.FieldCreatedAt:    {Type: field.TypeTime, Column: admin.FieldCreatedAt},
			admin.FieldUpdatedAt:    {Type: field.TypeTime, Column: admin.FieldUpdatedAt},
			admin.FieldIsEnable:     {Type: field.TypeBool, Column: admin.FieldIsEnable},
			admin.FieldDeletedAt:    {Type: field.TypeInt, Column: admin.FieldDeletedAt},
			admin.FieldUsername:     {Type: field.TypeString, Column: admin.FieldUsername},
			admin.FieldPassword:     {Type: field.TypeString, Column: admin.FieldPassword},
			admin.FieldWhitelistIps: {Type: field.TypeJSON, Column: admin.FieldWhitelistIps},
			admin.FieldDisplayName:  {Type: field.TypeString, Column: admin.FieldDisplayName},
			admin.FieldDepartmentID: {Type: field.TypeString, Column: admin.FieldDepartmentID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: department.FieldID,
			},
		},
		Type: "Department",
		Fields: map[string]*sqlgraph.FieldSpec{
			department.FieldCreatedAt: {Type: field.TypeTime, Column: department.FieldCreatedAt},
			department.FieldUpdatedAt: {Type: field.TypeTime, Column: department.FieldUpdatedAt},
			department.FieldDeletedAt: {Type: field.TypeInt, Column: department.FieldDeletedAt},
			department.FieldIsEnable:  {Type: field.TypeBool, Column: department.FieldIsEnable},
			department.FieldNameID:    {Type: field.TypeString, Column: department.FieldNameID},
			department.FieldName:      {Type: field.TypeString, Column: department.FieldName},
			department.FieldParentID:  {Type: field.TypeString, Column: department.FieldParentID},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: permission.FieldID,
			},
		},
		Type: "Permission",
		Fields: map[string]*sqlgraph.FieldSpec{
			permission.FieldCreatedAt: {Type: field.TypeTime, Column: permission.FieldCreatedAt},
			permission.FieldUpdatedAt: {Type: field.TypeTime, Column: permission.FieldUpdatedAt},
			permission.FieldGroup:     {Type: field.TypeString, Column: permission.FieldGroup},
			permission.FieldName:      {Type: field.TypeString, Column: permission.FieldName},
			permission.FieldKey:       {Type: field.TypeString, Column: permission.FieldKey},
			permission.FieldOrder:     {Type: field.TypeInt, Column: permission.FieldOrder},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   region.Table,
			Columns: region.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: region.FieldID,
			},
		},
		Type: "Region",
		Fields: map[string]*sqlgraph.FieldSpec{
			region.FieldCreatedAt: {Type: field.TypeTime, Column: region.FieldCreatedAt},
			region.FieldUpdatedAt: {Type: field.TypeTime, Column: region.FieldUpdatedAt},
			region.FieldDeletedAt: {Type: field.TypeInt, Column: region.FieldDeletedAt},
			region.FieldIsEnable:  {Type: field.TypeBool, Column: region.FieldIsEnable},
			region.FieldNameID:    {Type: field.TypeString, Column: region.FieldNameID},
			region.FieldName:      {Type: field.TypeString, Column: region.FieldName},
			region.FieldType:      {Type: field.TypeEnum, Column: region.FieldType},
			region.FieldParentID:  {Type: field.TypeString, Column: region.FieldParentID},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: role.FieldID,
			},
		},
		Type: "Role",
		Fields: map[string]*sqlgraph.FieldSpec{
			role.FieldCreatedAt:    {Type: field.TypeTime, Column: role.FieldCreatedAt},
			role.FieldUpdatedAt:    {Type: field.TypeTime, Column: role.FieldUpdatedAt},
			role.FieldIsEnable:     {Type: field.TypeBool, Column: role.FieldIsEnable},
			role.FieldDeletedAt:    {Type: field.TypeInt, Column: role.FieldDeletedAt},
			role.FieldName:         {Type: field.TypeString, Column: role.FieldName},
			role.FieldDescription:  {Type: field.TypeString, Column: role.FieldDescription},
			role.FieldOrder:        {Type: field.TypeInt, Column: role.FieldOrder},
			role.FieldIsChangeable: {Type: field.TypeBool, Column: role.FieldIsChangeable},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   route.Table,
			Columns: route.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: route.FieldID,
			},
		},
		Type: "Route",
		Fields: map[string]*sqlgraph.FieldSpec{
			route.FieldCreatedAt: {Type: field.TypeTime, Column: route.FieldCreatedAt},
			route.FieldUpdatedAt: {Type: field.TypeTime, Column: route.FieldUpdatedAt},
			route.FieldIsEnable:  {Type: field.TypeBool, Column: route.FieldIsEnable},
			route.FieldDeletedAt: {Type: field.TypeInt, Column: route.FieldDeletedAt},
			route.FieldParentID:  {Type: field.TypeString, Column: route.FieldParentID},
			route.FieldPath:      {Type: field.TypeString, Column: route.FieldPath},
			route.FieldComponent: {Type: field.TypeString, Column: route.FieldComponent},
			route.FieldRedirect:  {Type: field.TypeString, Column: route.FieldRedirect},
			route.FieldName:      {Type: field.TypeString, Column: route.FieldName},
			route.FieldOrder:     {Type: field.TypeInt, Column: route.FieldOrder},
			route.FieldType:      {Type: field.TypeEnum, Column: route.FieldType},
			route.FieldMeta:      {Type: field.TypeJSON, Column: route.FieldMeta},
		},
	}
	graph.MustAddE(
		"roles",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   admin.RolesTable,
			Columns: admin.RolesPrimaryKey,
			Bidi:    false,
		},
		"Admin",
		"Role",
	)
	graph.MustAddE(
		"department",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.DepartmentTable,
			Columns: []string{admin.DepartmentColumn},
			Bidi:    false,
		},
		"Admin",
		"Department",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
		},
		"Department",
		"Department",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
		},
		"Department",
		"Department",
	)
	graph.MustAddE(
		"members",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.MembersTable,
			Columns: []string{department.MembersColumn},
			Bidi:    false,
		},
		"Department",
		"Admin",
	)
	graph.MustAddE(
		"roles",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
		},
		"Permission",
		"Role",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   region.ParentTable,
			Columns: []string{region.ParentColumn},
			Bidi:    false,
		},
		"Region",
		"Region",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.ChildrenTable,
			Columns: []string{region.ChildrenColumn},
			Bidi:    false,
		},
		"Region",
		"Region",
	)
	graph.MustAddE(
		"admins",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.AdminsTable,
			Columns: role.AdminsPrimaryKey,
			Bidi:    false,
		},
		"Role",
		"Admin",
	)
	graph.MustAddE(
		"permissions",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
		},
		"Role",
		"Permission",
	)
	graph.MustAddE(
		"routes",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.RoutesTable,
			Columns: role.RoutesPrimaryKey,
			Bidi:    false,
		},
		"Role",
		"Route",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.ParentTable,
			Columns: []string{route.ParentColumn},
			Bidi:    false,
		},
		"Route",
		"Route",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.ChildrenTable,
			Columns: []string{route.ChildrenColumn},
			Bidi:    false,
		},
		"Route",
		"Route",
	)
	graph.MustAddE(
		"roles",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   route.RolesTable,
			Columns: route.RolesPrimaryKey,
			Bidi:    false,
		},
		"Route",
		"Role",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *AdminQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AdminQuery builder.
func (aq *AdminQuery) Filter() *AdminFilter {
	return &AdminFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *AdminMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AdminMutation builder.
func (m *AdminMutation) Filter() *AdminFilter {
	return &AdminFilter{config: m.config, predicateAdder: m}
}

// AdminFilter provides a generic filtering capability at runtime for AdminQuery.
type AdminFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AdminFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *AdminFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(admin.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *AdminFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(admin.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *AdminFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(admin.FieldUpdatedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *AdminFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(admin.FieldIsEnable))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *AdminFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(admin.FieldDeletedAt))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *AdminFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(admin.FieldUsername))
}

// WherePassword applies the entql string predicate on the password field.
func (f *AdminFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(admin.FieldPassword))
}

// WhereWhitelistIps applies the entql json.RawMessage predicate on the whitelist_ips field.
func (f *AdminFilter) WhereWhitelistIps(p entql.BytesP) {
	f.Where(p.Field(admin.FieldWhitelistIps))
}

// WhereDisplayName applies the entql string predicate on the display_name field.
func (f *AdminFilter) WhereDisplayName(p entql.StringP) {
	f.Where(p.Field(admin.FieldDisplayName))
}

// WhereDepartmentID applies the entql string predicate on the department_id field.
func (f *AdminFilter) WhereDepartmentID(p entql.StringP) {
	f.Where(p.Field(admin.FieldDepartmentID))
}

// WhereHasRoles applies a predicate to check if query has an edge roles.
func (f *AdminFilter) WhereHasRoles() {
	f.Where(entql.HasEdge("roles"))
}

// WhereHasRolesWith applies a predicate to check if query has an edge roles with a given conditions (other predicates).
func (f *AdminFilter) WhereHasRolesWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("roles", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasDepartment applies a predicate to check if query has an edge department.
func (f *AdminFilter) WhereHasDepartment() {
	f.Where(entql.HasEdge("department"))
}

// WhereHasDepartmentWith applies a predicate to check if query has an edge department with a given conditions (other predicates).
func (f *AdminFilter) WhereHasDepartmentWith(preds ...predicate.Department) {
	f.Where(entql.HasEdgeWith("department", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (dq *DepartmentQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DepartmentQuery builder.
func (dq *DepartmentQuery) Filter() *DepartmentFilter {
	return &DepartmentFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DepartmentMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DepartmentMutation builder.
func (m *DepartmentMutation) Filter() *DepartmentFilter {
	return &DepartmentFilter{config: m.config, predicateAdder: m}
}

// DepartmentFilter provides a generic filtering capability at runtime for DepartmentQuery.
type DepartmentFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DepartmentFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *DepartmentFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(department.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *DepartmentFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(department.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *DepartmentFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(department.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *DepartmentFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(department.FieldDeletedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *DepartmentFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(department.FieldIsEnable))
}

// WhereNameID applies the entql string predicate on the name_id field.
func (f *DepartmentFilter) WhereNameID(p entql.StringP) {
	f.Where(p.Field(department.FieldNameID))
}

// WhereName applies the entql string predicate on the name field.
func (f *DepartmentFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(department.FieldName))
}

// WhereParentID applies the entql string predicate on the parent_id field.
func (f *DepartmentFilter) WhereParentID(p entql.StringP) {
	f.Where(p.Field(department.FieldParentID))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *DepartmentFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *DepartmentFilter) WhereHasParentWith(preds ...predicate.Department) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *DepartmentFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *DepartmentFilter) WhereHasChildrenWith(preds ...predicate.Department) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasMembers applies a predicate to check if query has an edge members.
func (f *DepartmentFilter) WhereHasMembers() {
	f.Where(entql.HasEdge("members"))
}

// WhereHasMembersWith applies a predicate to check if query has an edge members with a given conditions (other predicates).
func (f *DepartmentFilter) WhereHasMembersWith(preds ...predicate.Admin) {
	f.Where(entql.HasEdgeWith("members", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (pq *PermissionQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PermissionQuery builder.
func (pq *PermissionQuery) Filter() *PermissionFilter {
	return &PermissionFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PermissionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PermissionMutation builder.
func (m *PermissionMutation) Filter() *PermissionFilter {
	return &PermissionFilter{config: m.config, predicateAdder: m}
}

// PermissionFilter provides a generic filtering capability at runtime for PermissionQuery.
type PermissionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PermissionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *PermissionFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(permission.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *PermissionFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(permission.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *PermissionFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(permission.FieldUpdatedAt))
}

// WhereGroup applies the entql string predicate on the group field.
func (f *PermissionFilter) WhereGroup(p entql.StringP) {
	f.Where(p.Field(permission.FieldGroup))
}

// WhereName applies the entql string predicate on the name field.
func (f *PermissionFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(permission.FieldName))
}

// WhereKey applies the entql string predicate on the key field.
func (f *PermissionFilter) WhereKey(p entql.StringP) {
	f.Where(p.Field(permission.FieldKey))
}

// WhereOrder applies the entql int predicate on the order field.
func (f *PermissionFilter) WhereOrder(p entql.IntP) {
	f.Where(p.Field(permission.FieldOrder))
}

// WhereHasRoles applies a predicate to check if query has an edge roles.
func (f *PermissionFilter) WhereHasRoles() {
	f.Where(entql.HasEdge("roles"))
}

// WhereHasRolesWith applies a predicate to check if query has an edge roles with a given conditions (other predicates).
func (f *PermissionFilter) WhereHasRolesWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("roles", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (rq *RegionQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RegionQuery builder.
func (rq *RegionQuery) Filter() *RegionFilter {
	return &RegionFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RegionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RegionMutation builder.
func (m *RegionMutation) Filter() *RegionFilter {
	return &RegionFilter{config: m.config, predicateAdder: m}
}

// RegionFilter provides a generic filtering capability at runtime for RegionQuery.
type RegionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RegionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *RegionFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(region.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *RegionFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(region.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *RegionFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(region.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *RegionFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(region.FieldDeletedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *RegionFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(region.FieldIsEnable))
}

// WhereNameID applies the entql string predicate on the name_id field.
func (f *RegionFilter) WhereNameID(p entql.StringP) {
	f.Where(p.Field(region.FieldNameID))
}

// WhereName applies the entql string predicate on the name field.
func (f *RegionFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(region.FieldName))
}

// WhereType applies the entql string predicate on the type field.
func (f *RegionFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(region.FieldType))
}

// WhereParentID applies the entql string predicate on the parent_id field.
func (f *RegionFilter) WhereParentID(p entql.StringP) {
	f.Where(p.Field(region.FieldParentID))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *RegionFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *RegionFilter) WhereHasParentWith(preds ...predicate.Region) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *RegionFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *RegionFilter) WhereHasChildrenWith(preds ...predicate.Region) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (rq *RoleQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RoleQuery builder.
func (rq *RoleQuery) Filter() *RoleFilter {
	return &RoleFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RoleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RoleMutation builder.
func (m *RoleMutation) Filter() *RoleFilter {
	return &RoleFilter{config: m.config, predicateAdder: m}
}

// RoleFilter provides a generic filtering capability at runtime for RoleQuery.
type RoleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RoleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *RoleFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(role.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *RoleFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *RoleFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldUpdatedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *RoleFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(role.FieldIsEnable))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *RoleFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(role.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *RoleFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(role.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *RoleFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(role.FieldDescription))
}

// WhereOrder applies the entql int predicate on the order field.
func (f *RoleFilter) WhereOrder(p entql.IntP) {
	f.Where(p.Field(role.FieldOrder))
}

// WhereIsChangeable applies the entql bool predicate on the is_changeable field.
func (f *RoleFilter) WhereIsChangeable(p entql.BoolP) {
	f.Where(p.Field(role.FieldIsChangeable))
}

// WhereHasAdmins applies a predicate to check if query has an edge admins.
func (f *RoleFilter) WhereHasAdmins() {
	f.Where(entql.HasEdge("admins"))
}

// WhereHasAdminsWith applies a predicate to check if query has an edge admins with a given conditions (other predicates).
func (f *RoleFilter) WhereHasAdminsWith(preds ...predicate.Admin) {
	f.Where(entql.HasEdgeWith("admins", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasPermissions applies a predicate to check if query has an edge permissions.
func (f *RoleFilter) WhereHasPermissions() {
	f.Where(entql.HasEdge("permissions"))
}

// WhereHasPermissionsWith applies a predicate to check if query has an edge permissions with a given conditions (other predicates).
func (f *RoleFilter) WhereHasPermissionsWith(preds ...predicate.Permission) {
	f.Where(entql.HasEdgeWith("permissions", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasRoutes applies a predicate to check if query has an edge routes.
func (f *RoleFilter) WhereHasRoutes() {
	f.Where(entql.HasEdge("routes"))
}

// WhereHasRoutesWith applies a predicate to check if query has an edge routes with a given conditions (other predicates).
func (f *RoleFilter) WhereHasRoutesWith(preds ...predicate.Route) {
	f.Where(entql.HasEdgeWith("routes", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (rq *RouteQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RouteQuery builder.
func (rq *RouteQuery) Filter() *RouteFilter {
	return &RouteFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RouteMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RouteMutation builder.
func (m *RouteMutation) Filter() *RouteFilter {
	return &RouteFilter{config: m.config, predicateAdder: m}
}

// RouteFilter provides a generic filtering capability at runtime for RouteQuery.
type RouteFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RouteFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *RouteFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(route.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *RouteFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(route.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *RouteFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(route.FieldUpdatedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *RouteFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(route.FieldIsEnable))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *RouteFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(route.FieldDeletedAt))
}

// WhereParentID applies the entql string predicate on the parent_id field.
func (f *RouteFilter) WhereParentID(p entql.StringP) {
	f.Where(p.Field(route.FieldParentID))
}

// WherePath applies the entql string predicate on the path field.
func (f *RouteFilter) WherePath(p entql.StringP) {
	f.Where(p.Field(route.FieldPath))
}

// WhereComponent applies the entql string predicate on the component field.
func (f *RouteFilter) WhereComponent(p entql.StringP) {
	f.Where(p.Field(route.FieldComponent))
}

// WhereRedirect applies the entql string predicate on the redirect field.
func (f *RouteFilter) WhereRedirect(p entql.StringP) {
	f.Where(p.Field(route.FieldRedirect))
}

// WhereName applies the entql string predicate on the name field.
func (f *RouteFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(route.FieldName))
}

// WhereOrder applies the entql int predicate on the order field.
func (f *RouteFilter) WhereOrder(p entql.IntP) {
	f.Where(p.Field(route.FieldOrder))
}

// WhereType applies the entql string predicate on the type field.
func (f *RouteFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(route.FieldType))
}

// WhereMeta applies the entql json.RawMessage predicate on the meta field.
func (f *RouteFilter) WhereMeta(p entql.BytesP) {
	f.Where(p.Field(route.FieldMeta))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *RouteFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *RouteFilter) WhereHasParentWith(preds ...predicate.Route) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *RouteFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *RouteFilter) WhereHasChildrenWith(preds ...predicate.Route) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasRoles applies a predicate to check if query has an edge roles.
func (f *RouteFilter) WhereHasRoles() {
	f.Where(entql.HasEdge("roles"))
}

// WhereHasRolesWith applies a predicate to check if query has an edge roles with a given conditions (other predicates).
func (f *RouteFilter) WhereHasRolesWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("roles", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
