package code_gen

var schemaTemplate = `package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type {{.Entity}} struct {
	ent.Schema
}

func ({{.Entity}}) Fields() []ent.Field {
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

func ({{.Entity}}) Edges() []ent.Edge {
	return nil
}

`
