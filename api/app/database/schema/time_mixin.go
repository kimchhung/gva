package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimetMixin struct {
	mixin.Schema
}

func (TimetMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").StructTag(`json:"createdAt,omitempty"`).Default(time.Now).Immutable(),
		field.Time("updated_at").StructTag(`json:"updatedAt,omitempty"`).Default(time.Now).UpdateDefault(time.Now).Immutable(),
	}
}
