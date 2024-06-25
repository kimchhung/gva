package schema

import (
	"github.com/gva/app/database/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Genre struct {
	ent.Schema
}

func (Genre) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NanoID{},
		mixins.TimeMixin{},
		mixins.GQLMixin{},
	}
}

func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").Values("comic"),
	}
}
