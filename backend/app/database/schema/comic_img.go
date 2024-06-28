package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pulid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ComicImg struct {
	ent.Schema
}

func (ComicImg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("CMI"),
		mixins.TimeMixin{},
	}
}

// Fields of the CloudflareImage.
func (ComicImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("b2key").Unique(),
		field.Int("height"),
		field.String("name"),
		field.Int64("optimized_size").StructTag(`json:"optimized"`),
		field.Int64("size").StructTag(`json:"s"`),
		field.Int("width").StructTag(`json:"w"`),
	}
}

func (ComicImg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chapter", ComicChapter.Type).
			Ref("imgs").Unique(),
	}
}
