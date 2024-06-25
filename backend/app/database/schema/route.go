package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Route struct {
	ent.Schema
}

func (Route) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NanoID{},
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Route) Fields() []ent.Field {
	return []ent.Field{
		field.String("parent_id").
			SchemaType(mixins.NanoSchemaType).
			DefaultFunc(mixins.NewNanoId).
			Optional().
			Nillable().
			StructTag(`json:"parentId,omitempty" rql:"filter,sort"`),

		field.String("path").
			StructTag(`rql:"filter,sort"`),

		field.String("component").
			StructTag(`rql:"filter,sort"`),

		field.String("redirect").
			Optional().
			Nillable().
			StructTag(`json:"redirect,omitempty" rql:"filter,sort"`),

		field.String("name").
			StructTag(`rql:"filter,sort"`),

		field.Int("order").
			StructTag(`rql:"filter,sort"`),

		field.Enum("type").
			Values("cata_log", "menu", "button", "external_link").
			Default("cata_log").
			StructTag(`rql:"filter,sort"`),

		field.JSON("meta", types.RouteMeta{}),
	}
}

func (Route) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Route.Type).From("parent").Unique().Field("parent_id"),
		edge.From("roles", Role.Type).Ref("routes"),
	}
}
