package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/extra/app/database/schema/mixins"
	"github.com/kimchhung/gva/extra/app/database/schema/softdelete"
)

type Route struct {
	ent.Schema
}

func (Route) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Route) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").Optional().Nillable().StructTag(`json:"parentId,omitempty"`),
		field.String("path"),
		field.String("component"),
		field.String("redirect"),
		field.String("name"),
		field.Int("id"),
		field.Int("type"),
		field.String("title"),
		field.JSON("meta", map[string]interface{}{}),
	}
}

func (Route) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Route.Type).From("parent").Unique().Field("parent_id"),
		edge.From("roles", Role.Type).Ref("routes"),
	}
}
