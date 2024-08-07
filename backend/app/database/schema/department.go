package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Department struct {
	ent.Schema
}

func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("dpm"),
		mixins.TimeMixin{},
		softdelete.SoftDeleteMixin{},
		mixins.IsEnableMixin{},
		mixins.GQLMixin{},
	}
}

func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_id").
			StructTag(`json:"nameId" rql:"column=name_id,filter,sort"`),

		field.String("name").
			StructTag(`json:"name" rql:"column=name,filter,sort"`),

		field.String("pid").
			GoType(pxid.ID("")).
			Optional().
			Nillable().
			StructTag(`json:"pid,omitempty" rql:"filter,sort"`),
	}
}

func (Department) Indexes() []ent.Index {
	return []ent.Index{
		softdelete.Index("name_id").Unique(),
	}
}

func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		// Department.Children && Department.Parent
		edge.To("children", Department.Type).
			From("parent").
			Unique().
			Field("pid"),

		// has many members
		edge.To("members", Admin.Type),
	}
}
