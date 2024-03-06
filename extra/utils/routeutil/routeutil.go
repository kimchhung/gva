package routeutil

import "github.com/kimchhung/gva/extra/internal/ent"

func GroupRouteToNested(flatRoutes []*ent.Route) (parentRoutes []*ent.Route) {
	parentMapRoute := make(map[int][]*ent.Route)

	for _, route := range flatRoutes {
		if route.ParentID != nil {
			parentMapRoute[*route.ParentID] = append(parentMapRoute[*route.ParentID], route)
		}
	}

	for _, route := range flatRoutes {
		route.Edges.Children = AppendRouteChildrens(route.ID, parentMapRoute)

		if route.ParentID == nil {
			parentRoutes = append(parentRoutes, route)
		}
	}

	return parentRoutes
}

func FlattenNestedRoutes(parentRoutes []*ent.Route) (flatRoutes []*ent.Route) {
	for _, p := range parentRoutes {
		flatRoutes = append(flatRoutes, p)

		if len(p.Edges.Children) > 0 {
			flatRoutes = append(flatRoutes, FlattenNestedRoutes(p.Edges.Children)...)
		}
	}

	return flatRoutes
}

func AppendRouteChildrens(parentID int, parentMapRoute map[int][]*ent.Route) (children []*ent.Route) {
	children, ok := parentMapRoute[parentID]
	if !ok {
		return nil
	}

	for _, child := range children {
		child.Edges.Children = AppendRouteChildrens(child.ID, parentMapRoute) // Recursively
	}

	return children
}
