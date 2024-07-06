// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/region"
	"github.com/gva/internal/ent/role"
	"github.com/gva/internal/ent/route"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AdminQuery) CollectFields(ctx context.Context, satisfies ...string) (*AdminQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AdminQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(admin.Columns))
		selectedFields = []string{admin.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, roleImplementors)...); err != nil {
				return err
			}
			a.WithNamedRoles(alias, func(wq *RoleQuery) {
				*wq = *query
			})

		case "department":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DepartmentClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, departmentImplementors)...); err != nil {
				return err
			}
			a.withDepartment = query
			if _, ok := fieldSeen[admin.FieldDepartmentID]; !ok {
				selectedFields = append(selectedFields, admin.FieldDepartmentID)
				fieldSeen[admin.FieldDepartmentID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[admin.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, admin.FieldCreatedAt)
				fieldSeen[admin.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[admin.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, admin.FieldUpdatedAt)
				fieldSeen[admin.FieldUpdatedAt] = struct{}{}
			}
		case "isEnable":
			if _, ok := fieldSeen[admin.FieldIsEnable]; !ok {
				selectedFields = append(selectedFields, admin.FieldIsEnable)
				fieldSeen[admin.FieldIsEnable] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[admin.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, admin.FieldDeletedAt)
				fieldSeen[admin.FieldDeletedAt] = struct{}{}
			}
		case "username":
			if _, ok := fieldSeen[admin.FieldUsername]; !ok {
				selectedFields = append(selectedFields, admin.FieldUsername)
				fieldSeen[admin.FieldUsername] = struct{}{}
			}
		case "whitelistIps":
			if _, ok := fieldSeen[admin.FieldWhitelistIps]; !ok {
				selectedFields = append(selectedFields, admin.FieldWhitelistIps)
				fieldSeen[admin.FieldWhitelistIps] = struct{}{}
			}
		case "displayName":
			if _, ok := fieldSeen[admin.FieldDisplayName]; !ok {
				selectedFields = append(selectedFields, admin.FieldDisplayName)
				fieldSeen[admin.FieldDisplayName] = struct{}{}
			}
		case "departmentID":
			if _, ok := fieldSeen[admin.FieldDepartmentID]; !ok {
				selectedFields = append(selectedFields, admin.FieldDepartmentID)
				fieldSeen[admin.FieldDepartmentID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type adminPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AdminPaginateOption
}

func newAdminPaginateArgs(rv map[string]any) *adminPaginateArgs {
	args := &adminPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &AdminOrder{Field: &AdminOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAdminOrder(order))
			}
		case *AdminOrder:
			if v != nil {
				args.opts = append(args.opts, WithAdminOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*AdminWhereInput); ok {
		args.opts = append(args.opts, WithAdminFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DepartmentQuery) CollectFields(ctx context.Context, satisfies ...string) (*DepartmentQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return d, nil
	}
	if err := d.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DepartmentQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(department.Columns))
		selectedFields = []string{department.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DepartmentClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, departmentImplementors)...); err != nil {
				return err
			}
			d.withParent = query
			if _, ok := fieldSeen[department.FieldParentID]; !ok {
				selectedFields = append(selectedFields, department.FieldParentID)
				fieldSeen[department.FieldParentID] = struct{}{}
			}

		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DepartmentClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, departmentImplementors)...); err != nil {
				return err
			}
			d.WithNamedChildren(alias, func(wq *DepartmentQuery) {
				*wq = *query
			})

		case "members":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AdminClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, adminImplementors)...); err != nil {
				return err
			}
			d.WithNamedMembers(alias, func(wq *AdminQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[department.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, department.FieldCreatedAt)
				fieldSeen[department.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[department.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, department.FieldUpdatedAt)
				fieldSeen[department.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[department.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, department.FieldDeletedAt)
				fieldSeen[department.FieldDeletedAt] = struct{}{}
			}
		case "isEnable":
			if _, ok := fieldSeen[department.FieldIsEnable]; !ok {
				selectedFields = append(selectedFields, department.FieldIsEnable)
				fieldSeen[department.FieldIsEnable] = struct{}{}
			}
		case "nameID":
			if _, ok := fieldSeen[department.FieldNameID]; !ok {
				selectedFields = append(selectedFields, department.FieldNameID)
				fieldSeen[department.FieldNameID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[department.FieldName]; !ok {
				selectedFields = append(selectedFields, department.FieldName)
				fieldSeen[department.FieldName] = struct{}{}
			}
		case "parentID":
			if _, ok := fieldSeen[department.FieldParentID]; !ok {
				selectedFields = append(selectedFields, department.FieldParentID)
				fieldSeen[department.FieldParentID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		d.Select(selectedFields...)
	}
	return nil
}

type departmentPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DepartmentPaginateOption
}

func newDepartmentPaginateArgs(rv map[string]any) *departmentPaginateArgs {
	args := &departmentPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &DepartmentOrder{Field: &DepartmentOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithDepartmentOrder(order))
			}
		case *DepartmentOrder:
			if v != nil {
				args.opts = append(args.opts, WithDepartmentOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*DepartmentWhereInput); ok {
		args.opts = append(args.opts, WithDepartmentFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PermissionQuery) CollectFields(ctx context.Context, satisfies ...string) (*PermissionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pe, nil
	}
	if err := pe.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pe *PermissionQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(permission.Columns))
		selectedFields = []string{permission.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: pe.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, roleImplementors)...); err != nil {
				return err
			}
			pe.WithNamedRoles(alias, func(wq *RoleQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[permission.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldCreatedAt)
				fieldSeen[permission.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[permission.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldUpdatedAt)
				fieldSeen[permission.FieldUpdatedAt] = struct{}{}
			}
		case "group":
			if _, ok := fieldSeen[permission.FieldGroup]; !ok {
				selectedFields = append(selectedFields, permission.FieldGroup)
				fieldSeen[permission.FieldGroup] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[permission.FieldName]; !ok {
				selectedFields = append(selectedFields, permission.FieldName)
				fieldSeen[permission.FieldName] = struct{}{}
			}
		case "key":
			if _, ok := fieldSeen[permission.FieldKey]; !ok {
				selectedFields = append(selectedFields, permission.FieldKey)
				fieldSeen[permission.FieldKey] = struct{}{}
			}
		case "order":
			if _, ok := fieldSeen[permission.FieldOrder]; !ok {
				selectedFields = append(selectedFields, permission.FieldOrder)
				fieldSeen[permission.FieldOrder] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pe.Select(selectedFields...)
	}
	return nil
}

type permissionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PermissionPaginateOption
}

func newPermissionPaginateArgs(rv map[string]any) *permissionPaginateArgs {
	args := &permissionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PermissionOrder{Field: &PermissionOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPermissionOrder(order))
			}
		case *PermissionOrder:
			if v != nil {
				args.opts = append(args.opts, WithPermissionOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PermissionWhereInput); ok {
		args.opts = append(args.opts, WithPermissionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RegionQuery) CollectFields(ctx context.Context, satisfies ...string) (*RegionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RegionQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(region.Columns))
		selectedFields = []string{region.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RegionClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, regionImplementors)...); err != nil {
				return err
			}
			r.withParent = query
			if _, ok := fieldSeen[region.FieldParentID]; !ok {
				selectedFields = append(selectedFields, region.FieldParentID)
				fieldSeen[region.FieldParentID] = struct{}{}
			}

		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RegionClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, regionImplementors)...); err != nil {
				return err
			}
			r.WithNamedChildren(alias, func(wq *RegionQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[region.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, region.FieldCreatedAt)
				fieldSeen[region.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[region.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, region.FieldUpdatedAt)
				fieldSeen[region.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[region.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, region.FieldDeletedAt)
				fieldSeen[region.FieldDeletedAt] = struct{}{}
			}
		case "isEnable":
			if _, ok := fieldSeen[region.FieldIsEnable]; !ok {
				selectedFields = append(selectedFields, region.FieldIsEnable)
				fieldSeen[region.FieldIsEnable] = struct{}{}
			}
		case "nameID":
			if _, ok := fieldSeen[region.FieldNameID]; !ok {
				selectedFields = append(selectedFields, region.FieldNameID)
				fieldSeen[region.FieldNameID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[region.FieldName]; !ok {
				selectedFields = append(selectedFields, region.FieldName)
				fieldSeen[region.FieldName] = struct{}{}
			}
		case "type":
			if _, ok := fieldSeen[region.FieldType]; !ok {
				selectedFields = append(selectedFields, region.FieldType)
				fieldSeen[region.FieldType] = struct{}{}
			}
		case "parentID":
			if _, ok := fieldSeen[region.FieldParentID]; !ok {
				selectedFields = append(selectedFields, region.FieldParentID)
				fieldSeen[region.FieldParentID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type regionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RegionPaginateOption
}

func newRegionPaginateArgs(rv map[string]any) *regionPaginateArgs {
	args := &regionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &RegionOrder{Field: &RegionOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithRegionOrder(order))
			}
		case *RegionOrder:
			if v != nil {
				args.opts = append(args.opts, WithRegionOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*RegionWhereInput); ok {
		args.opts = append(args.opts, WithRegionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*RoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoleQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(role.Columns))
		selectedFields = []string{role.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "admins":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AdminClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, adminImplementors)...); err != nil {
				return err
			}
			r.WithNamedAdmins(alias, func(wq *AdminQuery) {
				*wq = *query
			})

		case "permissions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PermissionClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, permissionImplementors)...); err != nil {
				return err
			}
			r.WithNamedPermissions(alias, func(wq *PermissionQuery) {
				*wq = *query
			})

		case "routes":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RouteClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, routeImplementors)...); err != nil {
				return err
			}
			r.WithNamedRoutes(alias, func(wq *RouteQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[role.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldCreatedAt)
				fieldSeen[role.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[role.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldUpdatedAt)
				fieldSeen[role.FieldUpdatedAt] = struct{}{}
			}
		case "isEnable":
			if _, ok := fieldSeen[role.FieldIsEnable]; !ok {
				selectedFields = append(selectedFields, role.FieldIsEnable)
				fieldSeen[role.FieldIsEnable] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[role.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldDeletedAt)
				fieldSeen[role.FieldDeletedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[role.FieldName]; !ok {
				selectedFields = append(selectedFields, role.FieldName)
				fieldSeen[role.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[role.FieldDescription]; !ok {
				selectedFields = append(selectedFields, role.FieldDescription)
				fieldSeen[role.FieldDescription] = struct{}{}
			}
		case "order":
			if _, ok := fieldSeen[role.FieldOrder]; !ok {
				selectedFields = append(selectedFields, role.FieldOrder)
				fieldSeen[role.FieldOrder] = struct{}{}
			}
		case "isChangeable":
			if _, ok := fieldSeen[role.FieldIsChangeable]; !ok {
				selectedFields = append(selectedFields, role.FieldIsChangeable)
				fieldSeen[role.FieldIsChangeable] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type rolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RolePaginateOption
}

func newRolePaginateArgs(rv map[string]any) *rolePaginateArgs {
	args := &rolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &RoleOrder{Field: &RoleOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithRoleOrder(order))
			}
		case *RoleOrder:
			if v != nil {
				args.opts = append(args.opts, WithRoleOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*RoleWhereInput); ok {
		args.opts = append(args.opts, WithRoleFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RouteQuery) CollectFields(ctx context.Context, satisfies ...string) (*RouteQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RouteQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(route.Columns))
		selectedFields = []string{route.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RouteClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, routeImplementors)...); err != nil {
				return err
			}
			r.withParent = query
			if _, ok := fieldSeen[route.FieldParentID]; !ok {
				selectedFields = append(selectedFields, route.FieldParentID)
				fieldSeen[route.FieldParentID] = struct{}{}
			}

		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RouteClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, routeImplementors)...); err != nil {
				return err
			}
			r.WithNamedChildren(alias, func(wq *RouteQuery) {
				*wq = *query
			})

		case "roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, roleImplementors)...); err != nil {
				return err
			}
			r.WithNamedRoles(alias, func(wq *RoleQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[route.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, route.FieldCreatedAt)
				fieldSeen[route.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[route.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, route.FieldUpdatedAt)
				fieldSeen[route.FieldUpdatedAt] = struct{}{}
			}
		case "isEnable":
			if _, ok := fieldSeen[route.FieldIsEnable]; !ok {
				selectedFields = append(selectedFields, route.FieldIsEnable)
				fieldSeen[route.FieldIsEnable] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[route.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, route.FieldDeletedAt)
				fieldSeen[route.FieldDeletedAt] = struct{}{}
			}
		case "parentID":
			if _, ok := fieldSeen[route.FieldParentID]; !ok {
				selectedFields = append(selectedFields, route.FieldParentID)
				fieldSeen[route.FieldParentID] = struct{}{}
			}
		case "path":
			if _, ok := fieldSeen[route.FieldPath]; !ok {
				selectedFields = append(selectedFields, route.FieldPath)
				fieldSeen[route.FieldPath] = struct{}{}
			}
		case "component":
			if _, ok := fieldSeen[route.FieldComponent]; !ok {
				selectedFields = append(selectedFields, route.FieldComponent)
				fieldSeen[route.FieldComponent] = struct{}{}
			}
		case "redirect":
			if _, ok := fieldSeen[route.FieldRedirect]; !ok {
				selectedFields = append(selectedFields, route.FieldRedirect)
				fieldSeen[route.FieldRedirect] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[route.FieldName]; !ok {
				selectedFields = append(selectedFields, route.FieldName)
				fieldSeen[route.FieldName] = struct{}{}
			}
		case "order":
			if _, ok := fieldSeen[route.FieldOrder]; !ok {
				selectedFields = append(selectedFields, route.FieldOrder)
				fieldSeen[route.FieldOrder] = struct{}{}
			}
		case "type":
			if _, ok := fieldSeen[route.FieldType]; !ok {
				selectedFields = append(selectedFields, route.FieldType)
				fieldSeen[route.FieldType] = struct{}{}
			}
		case "meta":
			if _, ok := fieldSeen[route.FieldMeta]; !ok {
				selectedFields = append(selectedFields, route.FieldMeta)
				fieldSeen[route.FieldMeta] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type routePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RoutePaginateOption
}

func newRoutePaginateArgs(rv map[string]any) *routePaginateArgs {
	args := &routePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &RouteOrder{Field: &RouteOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithRouteOrder(order))
			}
		case *RouteOrder:
			if v != nil {
				args.opts = append(args.opts, WithRouteOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*RouteWhereInput); ok {
		args.opts = append(args.opts, WithRouteFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
