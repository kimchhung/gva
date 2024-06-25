// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/gva/internal/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AdminQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AdminQueryRuleFunc func(context.Context, *ent.AdminQuery) error

// EvalQuery return f(ctx, q).
func (f AdminQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AdminQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AdminQuery", q)
}

// The AdminMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AdminMutationRuleFunc func(context.Context, *ent.AdminMutation) error

// EvalMutation calls f(ctx, m).
func (f AdminMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AdminMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AdminMutation", m)
}

// The ComicQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ComicQueryRuleFunc func(context.Context, *ent.ComicQuery) error

// EvalQuery return f(ctx, q).
func (f ComicQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ComicQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ComicQuery", q)
}

// The ComicMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ComicMutationRuleFunc func(context.Context, *ent.ComicMutation) error

// EvalMutation calls f(ctx, m).
func (f ComicMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ComicMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ComicMutation", m)
}

// The ComicChapterQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ComicChapterQueryRuleFunc func(context.Context, *ent.ComicChapterQuery) error

// EvalQuery return f(ctx, q).
func (f ComicChapterQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ComicChapterQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ComicChapterQuery", q)
}

// The ComicChapterMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ComicChapterMutationRuleFunc func(context.Context, *ent.ComicChapterMutation) error

// EvalMutation calls f(ctx, m).
func (f ComicChapterMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ComicChapterMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ComicChapterMutation", m)
}

// The ComicImgQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ComicImgQueryRuleFunc func(context.Context, *ent.ComicImgQuery) error

// EvalQuery return f(ctx, q).
func (f ComicImgQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ComicImgQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ComicImgQuery", q)
}

// The ComicImgMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ComicImgMutationRuleFunc func(context.Context, *ent.ComicImgMutation) error

// EvalMutation calls f(ctx, m).
func (f ComicImgMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ComicImgMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ComicImgMutation", m)
}

// The GenreQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GenreQueryRuleFunc func(context.Context, *ent.GenreQuery) error

// EvalQuery return f(ctx, q).
func (f GenreQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenreQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GenreQuery", q)
}

// The GenreMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GenreMutationRuleFunc func(context.Context, *ent.GenreMutation) error

// EvalMutation calls f(ctx, m).
func (f GenreMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GenreMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GenreMutation", m)
}

// The PermissionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PermissionQueryRuleFunc func(context.Context, *ent.PermissionQuery) error

// EvalQuery return f(ctx, q).
func (f PermissionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PermissionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PermissionQuery", q)
}

// The PermissionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PermissionMutationRuleFunc func(context.Context, *ent.PermissionMutation) error

// EvalMutation calls f(ctx, m).
func (f PermissionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PermissionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PermissionMutation", m)
}

// The RoleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RoleQueryRuleFunc func(context.Context, *ent.RoleQuery) error

// EvalQuery return f(ctx, q).
func (f RoleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RoleQuery", q)
}

// The RoleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RoleMutationRuleFunc func(context.Context, *ent.RoleMutation) error

// EvalMutation calls f(ctx, m).
func (f RoleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RoleMutation", m)
}

// The RouteQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RouteQueryRuleFunc func(context.Context, *ent.RouteQuery) error

// EvalQuery return f(ctx, q).
func (f RouteQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RouteQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RouteQuery", q)
}

// The RouteMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RouteMutationRuleFunc func(context.Context, *ent.RouteMutation) error

// EvalMutation calls f(ctx, m).
func (f RouteMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RouteMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RouteMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AdminQuery:
		return q.Filter(), nil
	case *ent.ComicQuery:
		return q.Filter(), nil
	case *ent.ComicChapterQuery:
		return q.Filter(), nil
	case *ent.ComicImgQuery:
		return q.Filter(), nil
	case *ent.GenreQuery:
		return q.Filter(), nil
	case *ent.PermissionQuery:
		return q.Filter(), nil
	case *ent.RoleQuery:
		return q.Filter(), nil
	case *ent.RouteQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AdminMutation:
		return m.Filter(), nil
	case *ent.ComicMutation:
		return m.Filter(), nil
	case *ent.ComicChapterMutation:
		return m.Filter(), nil
	case *ent.ComicImgMutation:
		return m.Filter(), nil
	case *ent.GenreMutation:
		return m.Filter(), nil
	case *ent.PermissionMutation:
		return m.Filter(), nil
	case *ent.RoleMutation:
		return m.Filter(), nil
	case *ent.RouteMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
