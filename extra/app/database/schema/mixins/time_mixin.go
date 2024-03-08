package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			StructTag(`json:"createdAt,omitempty" rql:"column=created_at,filter,sort"`).
			Default(time.Now),

		field.Time("updated_at").
			StructTag(`json:"updatedAt,omitempty"`).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (TimeMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"edges", rql:"-"`,
		},
		field.Annotation{
			StructTag: map[string]string{
				"id": `json:"id", rql:"filter,sort"`,
			},
		},
	}
}
