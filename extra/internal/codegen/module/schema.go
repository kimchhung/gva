package module_template

import "fmt"

var Schema = fmt.Sprintf(`package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type {{.EntityPascal}} struct {
	ent.Schema
}

func ({{.EntityPascal}}) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func ({{.EntityPascal}}) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(%sjson:"name,omitempty"%s),
	}
}

func ({{.EntityPascal}}) Edges() []ent.Edge {
	return nil
}


`, special, special)
