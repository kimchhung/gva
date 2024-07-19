package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gva/app/database/schema/mixins"
	"github.com/gva/app/database/schema/softdelete"
	"github.com/gva/app/database/schema/xid"
)

type MyTodo1 struct {
	ent.Schema
}

func (MyTodo1) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xid.MixinWithPrefix("mytodo1"),
		mixins.TimeMixin{},
		softdelete.SoftDeleteMixin{},
	}
}

func (MyTodo1) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name,omitempty"`),
	}
}

func (MyTodo1) Edges() []ent.Edge {
	return nil
}


