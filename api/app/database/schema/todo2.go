package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Todo2 struct {
	ent.Schema
}

func (Todo2) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimetMixin{},
	}
}

func (Todo2) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (Todo2) Edges() []ent.Edge {
	return nil
}


