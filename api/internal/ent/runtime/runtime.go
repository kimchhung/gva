// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/kimchhung/gva/app/database/schema"
	"github.com/kimchhung/gva/internal/ent/admin"
	"github.com/kimchhung/gva/internal/ent/role"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinHooks1 := adminMixin[1].Hooks()
	admin.Hooks[0] = adminMixinHooks1[0]
	adminMixinInters1 := adminMixin[1].Interceptors()
	admin.Interceptors[0] = adminMixinInters1[0]
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminMixinFields1 := adminMixin[1].Fields()
	_ = adminMixinFields1
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields0[1].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescDeletedAt is the schema descriptor for deleted_at field.
	adminDescDeletedAt := adminMixinFields1[0].Descriptor()
	// admin.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	admin.DefaultDeletedAt = adminDescDeletedAt.Default.(int)
	// adminDescIsActive is the schema descriptor for is_active field.
	adminDescIsActive := adminFields[3].Descriptor()
	// admin.DefaultIsActive holds the default value on creation for the is_active field.
	admin.DefaultIsActive = adminDescIsActive.Default.(bool)
	roleMixin := schema.Role{}.Mixin()
	roleMixinHooks1 := roleMixin[1].Hooks()
	role.Hooks[0] = roleMixinHooks1[0]
	roleMixinInters1 := roleMixin[1].Interceptors()
	role.Interceptors[0] = roleMixinInters1[0]
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields1 := roleMixin[1].Fields()
	_ = roleMixinFields1
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescDeletedAt is the schema descriptor for deleted_at field.
	roleDescDeletedAt := roleMixinFields1[0].Descriptor()
	// role.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	role.DefaultDeletedAt = roleDescDeletedAt.Default.(int)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
