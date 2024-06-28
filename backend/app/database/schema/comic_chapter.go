package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pulid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ComicChapter struct {
	ent.Schema
}

func (ComicChapter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("CMCT"),
		mixins.TimeMixin{},
	}
}

func (ComicChapter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("chapter"),
		field.String("title").Nillable().Optional(),
		field.String("volumn").Nillable().Optional(),
		field.String("lang"),
		field.Uint("up_count").StructTag(`json:"upCount"`).Default(0),
		field.Uint("down_count").StructTag(`json:"downCount"`).Default(0),
		field.Bool("is_last_chapter").StructTag(`json:"isLastChapter"`).Default(false),
	}
}

func (ComicChapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("imgs", ComicImg.Type),
		edge.From("comic", Comic.Type).
			Ref("chapters").Unique(),
	}
}
