package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/types"
	"github.com/gva/app/database/schema/xid"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Menu struct {
	ent.Schema
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xid.MixinWithPrefix("menu"),
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		softdelete.Index("path", "parent_id", "type").
			Unique().
			Annotations(
				entsql.IndexWhere("parent_id is null"),
			),
	}
}

func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("parent_id").
			GoType(xid.ID("")).
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
			Optional().
			Default(0).
			StructTag(`rql:"filter,sort"`),

		field.Enum("type").
			Values("cata_log", "menu", "button", "external_link").
			Default("cata_log").
			StructTag(`rql:"filter,sort"`),

		field.JSON("meta", types.MenuMeta{}),
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).From("parent").Unique().Field("parent_id"),
		edge.From("roles", Role.Type).Ref("routes"),
	}
}