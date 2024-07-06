package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/xid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xid.MixinWithPrefix("admin"),
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			StructTag(`json:"username" rql:"column=username,filter,sort"`).
			Unique(),

		field.String("password").
			Sensitive(),

		field.JSON("whitelist_ips", []string{}).
			StructTag(`json:"whitelistIps"`),

		field.String("display_name").
			StructTag(`json:"displayName,omitempty" rql:"filter,sort"`).
			Optional(),
	}
}

func (Admin) Indexes() []ent.Index {
	return []ent.Index{
		softdelete.Index("username").Unique(),
	}
}

func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}
