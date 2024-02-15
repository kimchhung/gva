package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimetMixin{},
	}
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
		field.String("display_name").StructTag(`json:"displayName,omitempty"`),
	}
}

func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}
