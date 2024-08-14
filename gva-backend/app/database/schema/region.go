package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"
)

type Region struct {
	ent.Schema
}

func (Region) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("regi"),
		mixins.TimeMixin{},
		softdelete.SoftDeleteMixin{},
		mixins.IsEnableMixin{},
		mixins.GQLMixin{},
	}
}

func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_id").
			StructTag(`json:"nameId" rql:"column=name_id,filter,sort"`),

		field.String("name").
			StructTag(`json:"name" rql:"column=name,filter,sort"`),

		field.Enum("type").
			Values("continent", "country", "city", "street", "any").
			StructTag(`json:"type" rql:"column=name,filter,sort"`),

		field.String("pid").
			GoType(pxid.ID("")).
			Optional().
			Nillable().
			StructTag(`json:"pid,omitempty" rql:"filter,sort"`),
	}
}

func (Region) Indexes() []ent.Index {
	return []ent.Index{
		softdelete.Index("name_id").Unique(),
	}
}

func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		// Department.Children && Department.Parent
		edge.To("children", Region.Type).
			From("parent").
			Unique().
			Field("pid"),
	}
}
