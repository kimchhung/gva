package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Permission struct {
	ent.Schema
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
		field.String("key").StructTag(`json:"key,omitempty"`),
		field.String("group").StructTag(`json:"group,omitempty"`),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions"),
	}
}
