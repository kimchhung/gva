// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/genre"
	"github.com/gva/internal/ent/manga"
	"github.com/gva/internal/ent/mangachapter"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/predicate"
	"github.com/gva/internal/ent/role"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 7)}
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
			department.FieldPid:       {Type: field.TypeString, Column: department.FieldPid},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   genre.Table,
			Columns: genre.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: genre.FieldID,
			},
		},
		Type: "Genre",
		Fields: map[string]*sqlgraph.FieldSpec{
			genre.FieldCreatedAt: {Type: field.TypeTime, Column: genre.FieldCreatedAt},
			genre.FieldUpdatedAt: {Type: field.TypeTime, Column: genre.FieldUpdatedAt},
			genre.FieldIsEnable:  {Type: field.TypeBool, Column: genre.FieldIsEnable},
			genre.FieldDeletedAt: {Type: field.TypeInt, Column: genre.FieldDeletedAt},
			genre.FieldName:      {Type: field.TypeString, Column: genre.FieldName},
			genre.FieldNameID:    {Type: field.TypeString, Column: genre.FieldNameID},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   manga.Table,
			Columns: manga.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: manga.FieldID,
			},
		},
		Type: "Manga",
		Fields: map[string]*sqlgraph.FieldSpec{
			manga.FieldCreatedAt:    {Type: field.TypeTime, Column: manga.FieldCreatedAt},
			manga.FieldUpdatedAt:    {Type: field.TypeTime, Column: manga.FieldUpdatedAt},
			manga.FieldIsEnable:     {Type: field.TypeBool, Column: manga.FieldIsEnable},
			manga.FieldDeletedAt:    {Type: field.TypeInt, Column: manga.FieldDeletedAt},
			manga.FieldNameID:       {Type: field.TypeString, Column: manga.FieldNameID},
			manga.FieldName:         {Type: field.TypeString, Column: manga.FieldName},
			manga.FieldDesc:         {Type: field.TypeString, Column: manga.FieldDesc},
			manga.FieldProdiver:     {Type: field.TypeString, Column: manga.FieldProdiver},
			manga.FieldThumbnailURL: {Type: field.TypeString, Column: manga.FieldThumbnailURL},
			manga.FieldAuthors:      {Type: field.TypeJSON, Column: manga.FieldAuthors},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   mangachapter.Table,
			Columns: mangachapter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mangachapter.FieldID,
			},
		},
		Type: "MangaChapter",
		Fields: map[string]*sqlgraph.FieldSpec{
			mangachapter.FieldCreatedAt:        {Type: field.TypeTime, Column: mangachapter.FieldCreatedAt},
			mangachapter.FieldUpdatedAt:        {Type: field.TypeTime, Column: mangachapter.FieldUpdatedAt},
			mangachapter.FieldMangaID:          {Type: field.TypeString, Column: mangachapter.FieldMangaID},
			mangachapter.FieldTitle:            {Type: field.TypeString, Column: mangachapter.FieldTitle},
			mangachapter.FieldImgURL:           {Type: field.TypeString, Column: mangachapter.FieldImgURL},
			mangachapter.FieldNumber:           {Type: field.TypeUint, Column: mangachapter.FieldNumber},
			mangachapter.FieldProviderName:     {Type: field.TypeString, Column: mangachapter.FieldProviderName},
			mangachapter.FieldChapterUpdatedAt: {Type: field.TypeTime, Column: mangachapter.FieldChapterUpdatedAt},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
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
			permission.FieldScope:     {Type: field.TypeString, Column: permission.FieldScope},
			permission.FieldType:      {Type: field.TypeEnum, Column: permission.FieldType},
			permission.FieldOrder:     {Type: field.TypeInt, Column: permission.FieldOrder},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
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
		"mangas",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   genre.MangasTable,
			Columns: genre.MangasPrimaryKey,
			Bidi:    false,
		},
		"Genre",
		"Manga",
	)
	graph.MustAddE(
		"chapters",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manga.ChaptersTable,
			Columns: []string{manga.ChaptersColumn},
			Bidi:    false,
		},
		"Manga",
		"MangaChapter",
	)
	graph.MustAddE(
		"genres",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   manga.GenresTable,
			Columns: manga.GenresPrimaryKey,
			Bidi:    false,
		},
		"Manga",
		"Genre",
	)
	graph.MustAddE(
		"manga",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mangachapter.MangaTable,
			Columns: []string{mangachapter.MangaColumn},
			Bidi:    false,
		},
		"MangaChapter",
		"Manga",
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

// WherePid applies the entql string predicate on the pid field.
func (f *DepartmentFilter) WherePid(p entql.StringP) {
	f.Where(p.Field(department.FieldPid))
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
func (gq *GenreQuery) addPredicate(pred func(s *sql.Selector)) {
	gq.predicates = append(gq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GenreQuery builder.
func (gq *GenreQuery) Filter() *GenreFilter {
	return &GenreFilter{config: gq.config, predicateAdder: gq}
}

// addPredicate implements the predicateAdder interface.
func (m *GenreMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GenreMutation builder.
func (m *GenreMutation) Filter() *GenreFilter {
	return &GenreFilter{config: m.config, predicateAdder: m}
}

// GenreFilter provides a generic filtering capability at runtime for GenreQuery.
type GenreFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GenreFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *GenreFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(genre.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *GenreFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(genre.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *GenreFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(genre.FieldUpdatedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *GenreFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(genre.FieldIsEnable))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *GenreFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(genre.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *GenreFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(genre.FieldName))
}

// WhereNameID applies the entql string predicate on the name_id field.
func (f *GenreFilter) WhereNameID(p entql.StringP) {
	f.Where(p.Field(genre.FieldNameID))
}

// WhereHasMangas applies a predicate to check if query has an edge mangas.
func (f *GenreFilter) WhereHasMangas() {
	f.Where(entql.HasEdge("mangas"))
}

// WhereHasMangasWith applies a predicate to check if query has an edge mangas with a given conditions (other predicates).
func (f *GenreFilter) WhereHasMangasWith(preds ...predicate.Manga) {
	f.Where(entql.HasEdgeWith("mangas", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (mq *MangaQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MangaQuery builder.
func (mq *MangaQuery) Filter() *MangaFilter {
	return &MangaFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MangaMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MangaMutation builder.
func (m *MangaMutation) Filter() *MangaFilter {
	return &MangaFilter{config: m.config, predicateAdder: m}
}

// MangaFilter provides a generic filtering capability at runtime for MangaQuery.
type MangaFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MangaFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *MangaFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(manga.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *MangaFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(manga.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *MangaFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(manga.FieldUpdatedAt))
}

// WhereIsEnable applies the entql bool predicate on the is_enable field.
func (f *MangaFilter) WhereIsEnable(p entql.BoolP) {
	f.Where(p.Field(manga.FieldIsEnable))
}

// WhereDeletedAt applies the entql int predicate on the deleted_at field.
func (f *MangaFilter) WhereDeletedAt(p entql.IntP) {
	f.Where(p.Field(manga.FieldDeletedAt))
}

// WhereNameID applies the entql string predicate on the name_id field.
func (f *MangaFilter) WhereNameID(p entql.StringP) {
	f.Where(p.Field(manga.FieldNameID))
}

// WhereName applies the entql string predicate on the name field.
func (f *MangaFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(manga.FieldName))
}

// WhereDesc applies the entql string predicate on the desc field.
func (f *MangaFilter) WhereDesc(p entql.StringP) {
	f.Where(p.Field(manga.FieldDesc))
}

// WhereProdiver applies the entql string predicate on the prodiver field.
func (f *MangaFilter) WhereProdiver(p entql.StringP) {
	f.Where(p.Field(manga.FieldProdiver))
}

// WhereThumbnailURL applies the entql string predicate on the thumbnail_url field.
func (f *MangaFilter) WhereThumbnailURL(p entql.StringP) {
	f.Where(p.Field(manga.FieldThumbnailURL))
}

// WhereAuthors applies the entql json.RawMessage predicate on the authors field.
func (f *MangaFilter) WhereAuthors(p entql.BytesP) {
	f.Where(p.Field(manga.FieldAuthors))
}

// WhereHasChapters applies a predicate to check if query has an edge chapters.
func (f *MangaFilter) WhereHasChapters() {
	f.Where(entql.HasEdge("chapters"))
}

// WhereHasChaptersWith applies a predicate to check if query has an edge chapters with a given conditions (other predicates).
func (f *MangaFilter) WhereHasChaptersWith(preds ...predicate.MangaChapter) {
	f.Where(entql.HasEdgeWith("chapters", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasGenres applies a predicate to check if query has an edge genres.
func (f *MangaFilter) WhereHasGenres() {
	f.Where(entql.HasEdge("genres"))
}

// WhereHasGenresWith applies a predicate to check if query has an edge genres with a given conditions (other predicates).
func (f *MangaFilter) WhereHasGenresWith(preds ...predicate.Genre) {
	f.Where(entql.HasEdgeWith("genres", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (mcq *MangaChapterQuery) addPredicate(pred func(s *sql.Selector)) {
	mcq.predicates = append(mcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MangaChapterQuery builder.
func (mcq *MangaChapterQuery) Filter() *MangaChapterFilter {
	return &MangaChapterFilter{config: mcq.config, predicateAdder: mcq}
}

// addPredicate implements the predicateAdder interface.
func (m *MangaChapterMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MangaChapterMutation builder.
func (m *MangaChapterMutation) Filter() *MangaChapterFilter {
	return &MangaChapterFilter{config: m.config, predicateAdder: m}
}

// MangaChapterFilter provides a generic filtering capability at runtime for MangaChapterQuery.
type MangaChapterFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MangaChapterFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *MangaChapterFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(mangachapter.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *MangaChapterFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(mangachapter.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *MangaChapterFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(mangachapter.FieldUpdatedAt))
}

// WhereMangaID applies the entql string predicate on the manga_id field.
func (f *MangaChapterFilter) WhereMangaID(p entql.StringP) {
	f.Where(p.Field(mangachapter.FieldMangaID))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *MangaChapterFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(mangachapter.FieldTitle))
}

// WhereImgURL applies the entql string predicate on the img_url field.
func (f *MangaChapterFilter) WhereImgURL(p entql.StringP) {
	f.Where(p.Field(mangachapter.FieldImgURL))
}

// WhereNumber applies the entql uint predicate on the number field.
func (f *MangaChapterFilter) WhereNumber(p entql.UintP) {
	f.Where(p.Field(mangachapter.FieldNumber))
}

// WhereProviderName applies the entql string predicate on the provider_name field.
func (f *MangaChapterFilter) WhereProviderName(p entql.StringP) {
	f.Where(p.Field(mangachapter.FieldProviderName))
}

// WhereChapterUpdatedAt applies the entql time.Time predicate on the chapter_updated_at field.
func (f *MangaChapterFilter) WhereChapterUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(mangachapter.FieldChapterUpdatedAt))
}

// WhereHasManga applies a predicate to check if query has an edge manga.
func (f *MangaChapterFilter) WhereHasManga() {
	f.Where(entql.HasEdge("manga"))
}

// WhereHasMangaWith applies a predicate to check if query has an edge manga with a given conditions (other predicates).
func (f *MangaChapterFilter) WhereHasMangaWith(preds ...predicate.Manga) {
	f.Where(entql.HasEdgeWith("manga", sqlgraph.WrapFunc(func(s *sql.Selector) {
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
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

// WhereScope applies the entql string predicate on the scope field.
func (f *PermissionFilter) WhereScope(p entql.StringP) {
	f.Where(p.Field(permission.FieldScope))
}

// WhereType applies the entql string predicate on the type field.
func (f *PermissionFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(permission.FieldType))
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
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
