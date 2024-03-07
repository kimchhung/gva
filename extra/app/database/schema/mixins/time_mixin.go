package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			StructTag(`json:"createdAt,omitempty" rql:"name=createdAt,column=created_at,filter,sort"`).
			Default(time.Now),

		field.Time("updated_at").
			StructTag(`json:"updatedAt,omitempty"`).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
