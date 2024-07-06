package softdelete_bak

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"

	// gen "github.com/gva/internal/ent"

	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const (
	deletedAtColumn = "deleted_at"
)

func Index(fields ...string) *index.Builder {
	fields = append(fields, deletedAtColumn)
	return index.Fields(fields...)
}

type SoftDeleteMixin struct {
	mixin.Schema
}

func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		// mysql doesn't support partial index yet!, uniq(col1,null_column) will be ignore so we need default value
		// if want nullable fields to support uniq index, required expr like this,  expr = "(CASE WHEN deleted_at IS NULL THEN CONCAT(user_id, payment_method_id, payment_identifier) END)"
		field.Int("deleted_at").Annotations(entsql.Default("0")).StructTag(`json:"-"`).Default(0),
	}
}

func (SoftDeleteMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		// intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		// 	// Skip soft-delete, means include soft-deleted entities.
		// 	if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
		// 		return nil
		// 	}

		// 	d.P(q)
		// 	return nil
		// }),
	}
}

func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		// hook.On(
		// 	func(next ent.Mutator) ent.Mutator {
		// 		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		// 			// Skip soft-delete, means delete the entity permanently.
		// 			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
		// 				return next.Mutate(ctx, m)
		// 			}
		// 			mx, ok := m.(interface {
		// 				SetOp(ent.Op)
		// 				Client() *gen.Client
		// 				SetDeletedAt(int)
		// 				WhereP(...func(*sql.Selector))
		// 			})
		// 			if !ok {
		// 				return nil, fmt.Errorf("unexpected mutation type %T", m)
		// 			}
		// 			d.P(mx)
		// 			mx.SetOp(ent.OpUpdate)
		// 			mx.SetDeletedAt(int(time.Now().Unix()))
		// 			return mx.Client().Mutate(ctx, m)
		// 		})
		// 	},
		// 	ent.OpDeleteOne|ent.OpDelete,
		// ),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldEQ(d.Fields()[0].Descriptor().Name, 0),
	)
}
