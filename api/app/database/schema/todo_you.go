package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type TodoYou struct {
	ent.Schema
}

func (TodoYou) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimetMixin{},
	}
}

func (TodoYou) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (TodoYou) Edges() []ent.Edge {
	return nil
}


