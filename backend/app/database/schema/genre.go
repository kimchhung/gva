package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pulid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Genre struct {
	ent.Schema
}

func (Genre) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("GEN"),
		mixins.TimeMixin{},
	}
}

func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").Values("comic"),
	}
}
