package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Permission struct {
	ent.Schema
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("perm"),
		mixins.TimeMixin{},
	}
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("group").
			StructTag(`json:"group,omitempty"`),

		field.String("name").
			StructTag(`json:"name,omitempty"`),

		field.String("scope").
			StructTag(`json:"scope,omitempty"`),

		field.Enum("type").
			Values("dynamic", "static").
			StructTag(`json:"key,omitempty"`).
			Optional().
			Default("dynamic"),

		field.Int("order").
			StructTag(`json:"order,omitempty"`).
			Optional().
			Default(0),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions"),
	}
}
