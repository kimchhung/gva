package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Menu struct {
	ent.Schema
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("menu"),
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Menu) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("pid").
			GoType(pxid.ID("")).
			Optional().
			Nillable().
			StructTag(`json:"pid,omitempty" rql:"filter,sort"`),

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
			Optional().
			Default(0).
			StructTag(`json:"order" rql:"filter,sort"`),

		field.Enum("type").
			Values("cata_log", "menu", "button", "external_link").
			Default("cata_log").
			StructTag(`rql:"filter,sort"`),

		field.JSON("meta", types.MenuMeta{}),
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).From("parent").Unique().Field("pid"),
		edge.From("roles", Role.Type).Ref("routes"),
	}
}
