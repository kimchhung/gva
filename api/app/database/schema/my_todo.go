package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type MyTodo struct {
	ent.Schema
}

func (MyTodo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimetMixin{},
	}
}

func (MyTodo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (MyTodo) Edges() []ent.Edge {
	return nil
}


