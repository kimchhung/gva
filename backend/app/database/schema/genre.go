package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/backend/app/database/schema/mixins"
)

type Genre struct {
	ent.Schema
}

func (Genre) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NanoID{},
		mixins.TimeMixin{},
	}
}

func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").Values("comic"),
	}
}
