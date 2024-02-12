package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Admin struct {
	ent.Schema
}

// Fields of the Article.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable(),
	}
}

// Edges of the User.
func (Admin) Edges() []ent.Edge {
	return nil
}
