// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/gva/app/database/schema"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/department"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/region"
	"github.com/gva/internal/ent/role"
	"github.com/gva/internal/ent/route"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinHooks3 := adminMixin[3].Hooks()
	admin.Hooks[0] = adminMixinHooks3[0]
	adminMixinInters3 := adminMixin[3].Interceptors()
	admin.Interceptors[0] = adminMixinInters3[0]
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminMixinFields1 := adminMixin[1].Fields()
	_ = adminMixinFields1
	adminMixinFields2 := adminMixin[2].Fields()
	_ = adminMixinFields2
	adminMixinFields3 := adminMixin[3].Fields()
	_ = adminMixinFields3
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields1[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields1[1].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescIsEnable is the schema descriptor for is_enable field.
	adminDescIsEnable := adminMixinFields2[0].Descriptor()
	// admin.DefaultIsEnable holds the default value on creation for the is_enable field.
	admin.DefaultIsEnable = adminDescIsEnable.Default.(bool)
	// adminDescDeletedAt is the schema descriptor for deleted_at field.
	adminDescDeletedAt := adminMixinFields3[0].Descriptor()
	// admin.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	admin.DefaultDeletedAt = adminDescDeletedAt.Default.(int)
	// adminDescID is the schema descriptor for id field.
	adminDescID := adminMixinFields0[0].Descriptor()
	// admin.DefaultID holds the default value on creation for the id field.
	admin.DefaultID = adminDescID.Default.(func() xid.ID)
	departmentMixin := schema.Department{}.Mixin()
	departmentMixinHooks2 := departmentMixin[2].Hooks()
	department.Hooks[0] = departmentMixinHooks2[0]
	departmentMixinInters2 := departmentMixin[2].Interceptors()
	department.Interceptors[0] = departmentMixinInters2[0]
	departmentMixinFields0 := departmentMixin[0].Fields()
	_ = departmentMixinFields0
	departmentMixinFields1 := departmentMixin[1].Fields()
	_ = departmentMixinFields1
	departmentMixinFields2 := departmentMixin[2].Fields()
	_ = departmentMixinFields2
	departmentMixinFields3 := departmentMixin[3].Fields()
	_ = departmentMixinFields3
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescCreatedAt is the schema descriptor for created_at field.
	departmentDescCreatedAt := departmentMixinFields1[0].Descriptor()
	// department.DefaultCreatedAt holds the default value on creation for the created_at field.
	department.DefaultCreatedAt = departmentDescCreatedAt.Default.(func() time.Time)
	// departmentDescUpdatedAt is the schema descriptor for updated_at field.
	departmentDescUpdatedAt := departmentMixinFields1[1].Descriptor()
	// department.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	department.DefaultUpdatedAt = departmentDescUpdatedAt.Default.(func() time.Time)
	// department.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	department.UpdateDefaultUpdatedAt = departmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// departmentDescDeletedAt is the schema descriptor for deleted_at field.
	departmentDescDeletedAt := departmentMixinFields2[0].Descriptor()
	// department.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	department.DefaultDeletedAt = departmentDescDeletedAt.Default.(int)
	// departmentDescIsEnable is the schema descriptor for is_enable field.
	departmentDescIsEnable := departmentMixinFields3[0].Descriptor()
	// department.DefaultIsEnable holds the default value on creation for the is_enable field.
	department.DefaultIsEnable = departmentDescIsEnable.Default.(bool)
	// departmentDescID is the schema descriptor for id field.
	departmentDescID := departmentMixinFields0[0].Descriptor()
	// department.DefaultID holds the default value on creation for the id field.
	department.DefaultID = departmentDescID.Default.(func() xid.ID)
	permissionMixin := schema.Permission{}.Mixin()
	permissionMixinFields0 := permissionMixin[0].Fields()
	_ = permissionMixinFields0
	permissionMixinFields1 := permissionMixin[1].Fields()
	_ = permissionMixinFields1
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionMixinFields1[0].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionMixinFields1[1].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(func() time.Time)
	// permission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permission.UpdateDefaultUpdatedAt = permissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionMixinFields0[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() xid.ID)
	regionMixin := schema.Region{}.Mixin()
	regionMixinHooks2 := regionMixin[2].Hooks()
	region.Hooks[0] = regionMixinHooks2[0]
	regionMixinInters2 := regionMixin[2].Interceptors()
	region.Interceptors[0] = regionMixinInters2[0]
	regionMixinFields0 := regionMixin[0].Fields()
	_ = regionMixinFields0
	regionMixinFields1 := regionMixin[1].Fields()
	_ = regionMixinFields1
	regionMixinFields2 := regionMixin[2].Fields()
	_ = regionMixinFields2
	regionMixinFields3 := regionMixin[3].Fields()
	_ = regionMixinFields3
	regionFields := schema.Region{}.Fields()
	_ = regionFields
	// regionDescCreatedAt is the schema descriptor for created_at field.
	regionDescCreatedAt := regionMixinFields1[0].Descriptor()
	// region.DefaultCreatedAt holds the default value on creation for the created_at field.
	region.DefaultCreatedAt = regionDescCreatedAt.Default.(func() time.Time)
	// regionDescUpdatedAt is the schema descriptor for updated_at field.
	regionDescUpdatedAt := regionMixinFields1[1].Descriptor()
	// region.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	region.DefaultUpdatedAt = regionDescUpdatedAt.Default.(func() time.Time)
	// region.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	region.UpdateDefaultUpdatedAt = regionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// regionDescDeletedAt is the schema descriptor for deleted_at field.
	regionDescDeletedAt := regionMixinFields2[0].Descriptor()
	// region.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	region.DefaultDeletedAt = regionDescDeletedAt.Default.(int)
	// regionDescIsEnable is the schema descriptor for is_enable field.
	regionDescIsEnable := regionMixinFields3[0].Descriptor()
	// region.DefaultIsEnable holds the default value on creation for the is_enable field.
	region.DefaultIsEnable = regionDescIsEnable.Default.(bool)
	// regionDescID is the schema descriptor for id field.
	regionDescID := regionMixinFields0[0].Descriptor()
	// region.DefaultID holds the default value on creation for the id field.
	region.DefaultID = regionDescID.Default.(func() xid.ID)
	roleMixin := schema.Role{}.Mixin()
	roleMixinHooks3 := roleMixin[3].Hooks()
	role.Hooks[0] = roleMixinHooks3[0]
	roleMixinInters3 := roleMixin[3].Interceptors()
	role.Interceptors[0] = roleMixinInters3[0]
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields1 := roleMixin[1].Fields()
	_ = roleMixinFields1
	roleMixinFields2 := roleMixin[2].Fields()
	_ = roleMixinFields2
	roleMixinFields3 := roleMixin[3].Fields()
	_ = roleMixinFields3
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields1[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields1[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescIsEnable is the schema descriptor for is_enable field.
	roleDescIsEnable := roleMixinFields2[0].Descriptor()
	// role.DefaultIsEnable holds the default value on creation for the is_enable field.
	role.DefaultIsEnable = roleDescIsEnable.Default.(bool)
	// roleDescDeletedAt is the schema descriptor for deleted_at field.
	roleDescDeletedAt := roleMixinFields3[0].Descriptor()
	// role.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	role.DefaultDeletedAt = roleDescDeletedAt.Default.(int)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() xid.ID)
	routeMixin := schema.Route{}.Mixin()
	routeMixinHooks3 := routeMixin[3].Hooks()
	route.Hooks[0] = routeMixinHooks3[0]
	routeMixinInters3 := routeMixin[3].Interceptors()
	route.Interceptors[0] = routeMixinInters3[0]
	routeMixinFields0 := routeMixin[0].Fields()
	_ = routeMixinFields0
	routeMixinFields1 := routeMixin[1].Fields()
	_ = routeMixinFields1
	routeMixinFields2 := routeMixin[2].Fields()
	_ = routeMixinFields2
	routeMixinFields3 := routeMixin[3].Fields()
	_ = routeMixinFields3
	routeFields := schema.Route{}.Fields()
	_ = routeFields
	// routeDescCreatedAt is the schema descriptor for created_at field.
	routeDescCreatedAt := routeMixinFields1[0].Descriptor()
	// route.DefaultCreatedAt holds the default value on creation for the created_at field.
	route.DefaultCreatedAt = routeDescCreatedAt.Default.(func() time.Time)
	// routeDescUpdatedAt is the schema descriptor for updated_at field.
	routeDescUpdatedAt := routeMixinFields1[1].Descriptor()
	// route.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	route.DefaultUpdatedAt = routeDescUpdatedAt.Default.(func() time.Time)
	// route.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	route.UpdateDefaultUpdatedAt = routeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// routeDescIsEnable is the schema descriptor for is_enable field.
	routeDescIsEnable := routeMixinFields2[0].Descriptor()
	// route.DefaultIsEnable holds the default value on creation for the is_enable field.
	route.DefaultIsEnable = routeDescIsEnable.Default.(bool)
	// routeDescDeletedAt is the schema descriptor for deleted_at field.
	routeDescDeletedAt := routeMixinFields3[0].Descriptor()
	// route.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	route.DefaultDeletedAt = routeDescDeletedAt.Default.(int)
	// routeDescOrder is the schema descriptor for order field.
	routeDescOrder := routeFields[5].Descriptor()
	// route.DefaultOrder holds the default value on creation for the order field.
	route.DefaultOrder = routeDescOrder.Default.(int)
	// routeDescID is the schema descriptor for id field.
	routeDescID := routeMixinFields0[0].Descriptor()
	// route.DefaultID holds the default value on creation for the id field.
	route.DefaultID = routeDescID.Default.(func() xid.ID)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
