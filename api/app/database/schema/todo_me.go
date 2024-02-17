package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type TodoMe struct {
	ent.Schema
}

func (TodoMe) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimetMixin{},
	}
}

func (TodoMe) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (TodoMe) Edges() []ent.Edge {
	return nil
}


