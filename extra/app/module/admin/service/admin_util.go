package service

import "github.com/kimchhung/gva/extra/internal/ent"

func groupRouteToNested(flatRoutes []*ent.Route) (parentRoutes []*ent.Route) {
	parentMapRoute := make(map[int][]*ent.Route)

	for _, route := range flatRoutes {
		if route.ParentID != nil {
			parentMapRoute[*route.ParentID] = append(parentMapRoute[*route.ParentID], route)
		}
	}

	for _, route := range flatRoutes {
		route.Edges.Children = appendRouteChildrens(route.ID, parentMapRoute)

		if route.ParentID == nil {
			parentRoutes = append(parentRoutes, route)
		}
	}

	return parentRoutes
}

// func flattenNestedRoutes(parentRoutes []*ent.Route) (flatRoutes []*ent.Route) {
// 	for _, p := range parentRoutes {
// 		if len(p.Edges.Children) > 0 {
// 			flatRoutes = append(flatRoutes, flattenNestedRoutes(p.Edges.Children)...)
// 		} else {
// 			flatRoutes = append(flatRoutes, p)
// 		}
// 	}

// 	return flatRoutes
// }

func appendRouteChildrens(parentID int, parentMapRoute map[int][]*ent.Route) (children []*ent.Route) {
	children, ok := parentMapRoute[parentID]
	if !ok {
		return nil
	}

	for _, child := range children {
		child.Edges.Children = appendRouteChildrens(child.ID, parentMapRoute) // Recursively
	}

	return children
}
