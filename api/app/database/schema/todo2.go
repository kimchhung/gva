package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Todo2 struct {
	ent.Schema
}

func (Todo2) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable(),
	}
}

func (Todo2) Edges() []ent.Edge {
	return nil
}

