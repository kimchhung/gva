package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/xid"
)

type MyTodo struct {
	ent.Schema
}

func (MyTodo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xid.MixinWithPrefix("mytodo"),
		mixins.TimeMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (MyTodo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (MyTodo) Edges() []ent.Edge {
	return nil
}


