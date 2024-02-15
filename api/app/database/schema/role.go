package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimetMixin{},
	}
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admins", Admin.Type).Ref("roles"),
		edge.To("permissions", Permission.Type),
	}
}
