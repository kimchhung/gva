package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/kimchhung/gva/backend/app/database/schema/mixins"
	"github.com/kimchhung/gva/backend/app/database/schema/softdelete"
)

type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NanoID{},
		mixins.TimeMixin{},
		mixins.IsEnableMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),

		field.String("description").StructTag(`json:"description,omitempty"`),

		field.Int("order").StructTag(`json:"order,omitempty"`),

		field.Bool("is_changeable").StructTag(`json:"isChangeable,omitempty"`),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admins", Admin.Type).Ref("roles"),
		edge.To("permissions", Permission.Type),
		edge.To("routes", Route.Type),
	}
}
