package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/app/database/schema/softdelete"
)

type Todo struct {
	ent.Schema
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pxid.MixinWithPrefix("todo"),
		mixins.TimeMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (Todo) Edges() []ent.Edge {
	return nil
}
