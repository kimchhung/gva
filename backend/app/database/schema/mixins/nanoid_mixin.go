package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type NanoID struct {
	mixin.Schema
}

func (NanoID) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").SchemaType(NanoSchemaType),
	}
}

var NanoSchemaType = map[string]string{
	dialect.MySQL: "VARCHAR(21)",
}
