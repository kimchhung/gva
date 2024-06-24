package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type NanoID struct {
	mixin.Schema
}

func (NanoID) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").SchemaType(NanoSchemaType).DefaultFunc(NewNanoId),
	}
}

func NewNanoId() string {
	id, err := gonanoid.New(21)
	if err != nil {
		panic(err)
	}
	return id
}

var NanoSchemaType = map[string]string{
	dialect.MySQL: "VARCHAR(21)",
}
