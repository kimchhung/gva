// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/comic"
	"github.com/gva/internal/ent/comicchapter"
	"github.com/gva/internal/ent/comicimg"
	"github.com/gva/internal/ent/genre"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/role"
	"github.com/gva/internal/ent/route"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
	IsNode()
}

var adminImplementors = []string{"Admin", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Admin) IsNode() {}

var comicImplementors = []string{"Comic", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Comic) IsNode() {}

var comicchapterImplementors = []string{"ComicChapter", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ComicChapter) IsNode() {}

var comicimgImplementors = []string{"ComicImg", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ComicImg) IsNode() {}

var genreImplementors = []string{"Genre", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Genre) IsNode() {}

var permissionImplementors = []string{"Permission", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Permission) IsNode() {}

var roleImplementors = []string{"Role", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Role) IsNode() {}

var routeImplementors = []string{"Route", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Route) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, pulid.ID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, pulid.ID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, pulid.ID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id pulid.ID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id pulid.ID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id pulid.ID) (Noder, error) {
	switch table {
	case admin.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Admin.Query().
			Where(admin.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, adminImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case comic.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Comic.Query().
			Where(comic.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, comicImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case comicchapter.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.ComicChapter.Query().
			Where(comicchapter.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, comicchapterImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case comicimg.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.ComicImg.Query().
			Where(comicimg.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, comicimgImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case genre.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Genre.Query().
			Where(genre.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, genreImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case permission.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Permission.Query().
			Where(permission.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, permissionImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case role.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Role.Query().
			Where(role.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, roleImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case route.Table:
		var uid pulid.ID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Route.Query().
			Where(route.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, routeImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []pulid.ID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]pulid.ID)
	id2idx := make(map[pulid.ID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []pulid.ID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[pulid.ID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case admin.Table:
		query := c.Admin.Query().
			Where(admin.IDIn(ids...))
		query, err := query.CollectFields(ctx, adminImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case comic.Table:
		query := c.Comic.Query().
			Where(comic.IDIn(ids...))
		query, err := query.CollectFields(ctx, comicImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case comicchapter.Table:
		query := c.ComicChapter.Query().
			Where(comicchapter.IDIn(ids...))
		query, err := query.CollectFields(ctx, comicchapterImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case comicimg.Table:
		query := c.ComicImg.Query().
			Where(comicimg.IDIn(ids...))
		query, err := query.CollectFields(ctx, comicimgImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case genre.Table:
		query := c.Genre.Query().
			Where(genre.IDIn(ids...))
		query, err := query.CollectFields(ctx, genreImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case permission.Table:
		query := c.Permission.Query().
			Where(permission.IDIn(ids...))
		query, err := query.CollectFields(ctx, permissionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case role.Table:
		query := c.Role.Query().
			Where(role.IDIn(ids...))
		query, err := query.CollectFields(ctx, roleImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case route.Table:
		query := c.Route.Query().
			Where(route.IDIn(ids...))
		query, err := query.CollectFields(ctx, routeImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
