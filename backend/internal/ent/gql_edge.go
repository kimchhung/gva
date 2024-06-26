// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Admin) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryRoles().All(ctx)
	}
	return result, err
}

func (c *Comic) Chapters(ctx context.Context) (result []*ComicChapter, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedChapters(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ChaptersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryChapters().All(ctx)
	}
	return result, err
}

func (c *Comic) LastChapter(ctx context.Context) (*ComicChapter, error) {
	result, err := c.Edges.LastChapterOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryLastChapter().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Comic) FinalChapter(ctx context.Context) (*ComicChapter, error) {
	result, err := c.Edges.FinalChapterOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryFinalChapter().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cc *ComicChapter) Imgs(ctx context.Context) (result []*ComicImg, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cc.NamedImgs(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cc.Edges.ImgsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cc.QueryImgs().All(ctx)
	}
	return result, err
}

func (cc *ComicChapter) Comic(ctx context.Context) (*Comic, error) {
	result, err := cc.Edges.ComicOrErr()
	if IsNotLoaded(err) {
		result, err = cc.QueryComic().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ci *ComicImg) Chapter(ctx context.Context) (*ComicChapter, error) {
	result, err := ci.Edges.ChapterOrErr()
	if IsNotLoaded(err) {
		result, err = ci.QueryChapter().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pe *Permission) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryRoles().All(ctx)
	}
	return result, err
}

func (r *Role) Admins(ctx context.Context) (result []*Admin, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedAdmins(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.AdminsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryAdmins().All(ctx)
	}
	return result, err
}

func (r *Role) Permissions(ctx context.Context) (result []*Permission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedPermissions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.PermissionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryPermissions().All(ctx)
	}
	return result, err
}

func (r *Role) Routes(ctx context.Context) (result []*Route, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedRoutes(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.RoutesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryRoutes().All(ctx)
	}
	return result, err
}

func (r *Route) Parent(ctx context.Context) (*Route, error) {
	result, err := r.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Route) Children(ctx context.Context) (result []*Route, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryChildren().All(ctx)
	}
	return result, err
}

func (r *Route) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryRoles().All(ctx)
	}
	return result, err
}
