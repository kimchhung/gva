package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Manga struct {
	ent.Schema
}

func (Manga) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("mga"),
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Manga) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_id").
			StructTag(`json:"name_id" rql:"column=name_id,filter,sort"`),

		field.String("name").
			StructTag(`json:"name" rql:"column=name,filter,sort"`),

		field.String("desc").
			StructTag(`json:"desc"`),

		field.String("prodiver").
			StructTag(`json:"provider"`),

		field.String("thumbnail_url").
			StructTag(`json:"thumbnailUrl"`),

		field.JSON("authors", []string{}).
			StructTag(`json:"authors"`),
	}
}

func (Manga) Indexes() []ent.Index {
	return []ent.Index{
		softdelete.Index("name", "name_id").Unique(),
	}
}

func (Manga) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chapters", MangaChapter.Type),
		edge.From("genres", Genre.Type).Ref("mangas"),
	}
}
