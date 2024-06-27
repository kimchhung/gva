package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Comic struct {
	ent.Schema
}

func (Comic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NanoID{},
		mixins.TimeMixin{},
	}
}

func (Comic) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("chapter"),
		field.String("title"),
		field.String("slug"),
		field.JSON("covers", []types.CoverImg{}),
		field.String("status"),
		field.Bool("isTranslateCompleted").
			StructTag("isTranslateCompleted").
			Default(false),
		field.Uint("up_count").StructTag(`json:"upCount"`).Default(0),

		field.String("final_chapter_id").
			SchemaType(mixins.NanoSchemaType).
			Optional().Nillable(),

		field.String("last_chapter_id").
			SchemaType(mixins.NanoSchemaType).
			Optional().Nillable(),
	}
}

func (Comic) Edges() []ent.Edge {
	return []ent.Edge{
		// Edge to all chapters
		edge.To("chapters", ComicChapter.Type),

		// Edge to last chapter
		edge.To("last_chapter", ComicChapter.Type).
			Unique().Field("last_chapter_id").StructTag(`json:"lastChapter"`),

		// Edge to final chapter
		edge.To("final_chapter", ComicChapter.Type).
			Unique().Field("final_chapter_id").StructTag(`json:"finalChapter"`),
	}
}
