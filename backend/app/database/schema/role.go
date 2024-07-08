package schema

import (
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/xid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xid.MixinWithPrefix("role"),
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
		edge.To("routes", Menu.Type),
	}
}
