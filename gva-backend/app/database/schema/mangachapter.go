package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
)

type MangaChapter struct {
	ent.Schema
}

func (MangaChapter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("mcr"),
		mixins.TimeMixin{},
	}
}

func (MangaChapter) Fields() []ent.Field {
	return []ent.Field{
		field.String("manga_id").
			GoType(pxid.ID("")).NotEmpty().
			StructTag(`json:"mangaId,omitempty" rql:"filter,sort"`),

		field.String("title").
			StructTag(`json:"title" rql:"filter,sort"`),

		field.String("img_url").
			StructTag(`json:"imgUrl"`),

		field.Uint("number").
			StructTag(`json:"number" rql:"filter,sort"`),

		field.String("provider_name").
			StructTag(`json:"providerName"`),

		field.Time("chapter_updated_at").
			StructTag(`json:"chapterUpdatedAt" rql:"filter,sort"`),
	}
}

func (MangaChapter) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider_name", "manga_id", "number").Unique(),
	}
}

func (MangaChapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("manga", Manga.Type).
			Ref("chapters").
			Unique().
			Field("manga_id").Required(),
	}
}
