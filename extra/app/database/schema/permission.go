package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/extra/app/database/schema/mixins"
)

type Permission struct {
	ent.Schema
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("group").StructTag(`json:"group,omitempty"`),
		field.String("name").StructTag(`json:"name,omitempty"`),
		field.String("key").StructTag(`json:"key,omitempty"`),
		field.Int("order").StructTag(`json:"order,omitempty"`),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions"),
	}
}
