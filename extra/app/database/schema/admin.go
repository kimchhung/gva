package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/extra/app/database/schema/mixins"
	"github.com/kimchhung/gva/extra/app/database/schema/softdelete"
)

type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`json:"id" rql:"column=id,name=id,filter,sort"`),

		field.String("username").
			StructTag(`json:"username" rql:"column=username,name=username,filter,sort"`).
			Unique(),

		field.String("password").
			Sensitive(),

		field.JSON("whitelist_ips", []string{}).
			StructTag(`json:"-"`),

		field.String("display_name").
			StructTag(`json:"displayName,omitempty" rql:"column=display_name,name=displayName,filter,sort"`).
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
