package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"
)

type Genre struct {
	ent.Schema
}

func (Genre) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("magr"),
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			StructTag(`json:"name" rql:"column=name,filter,sort"`),

		field.String("name_id").
			StructTag(`json:"name" rql:"column=name,filter,sort"`),
	}
}

func (Genre) Indexes() []ent.Index {
	return []ent.Index{
		softdelete.Index("name_id").Unique(),
	}
}

func (Genre) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mangas", Manga.Type),
	}
}
