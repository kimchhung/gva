package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type IsEnableMixin struct {
	mixin.Schema
}

func (IsEnableMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			Bool("is_enable").
			StructTag(`json:"isEnable"`).
			Default(true),
	}
}
