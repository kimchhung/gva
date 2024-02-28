package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/app/database/schema/softdelete"
)

type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password").Sensitive(),
		field.JSON("whitelist_ips", []string{}).StructTag(`json:"-"`),
		field.Bool("is_active").Default(true).StructTag(`json:"isActive"`),
		field.String("display_name").StructTag(`json:"displayName,omitempty"`).Optional(),
	}
}

func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}
