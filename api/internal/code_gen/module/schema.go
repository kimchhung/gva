package module_template

import "fmt"

var Schema = fmt.Sprintf(`package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type {{.Entity}} struct {
	ent.Schema
}

func ({{.Entity}}) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreatedAtMixin{},
		UpdatedAtMixin{},
		DeletedAtMixin{},
	}
}

func ({{.Entity}}) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(%sjson:"name,omitempty"%s),
	}
}

func ({{.Entity}}) Edges() []ent.Edge {
	return nil
}


`, special, special)
