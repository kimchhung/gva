package routeutil

import (
	"fmt"
	"slices"

	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/color"

	"github.com/gosuri/uitable"
	"github.com/labstack/echo/v4"
)

func GroupRouteToNested(flatRoutes []*ent.Route) (parentRoutes []*ent.Route) {
	parentMapRoute := make(map[xid.ID][]*ent.Route)

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

	slices.SortFunc(parentRoutes, func(a *ent.Route, b *ent.Route) int {
		return a.Order - b.Order
	})

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

func AppendRouteChildrens(parentID xid.ID, parentMapRoute map[xid.ID][]*ent.Route) (children []*ent.Route) {
	children, ok := parentMapRoute[parentID]
	if !ok {
		return nil
	}

	for _, child := range children {
		child.Edges.Children = AppendRouteChildrens(child.ID, parentMapRoute) // Recursively
	}

	return children
}

func PrintRoutes(routes []*echo.Route) {

	// Create a new table
	table := uitable.New()

	// Set the table headers

	table.AddRow("Method", "Path", "Name")
	for _, r := range routes {
		table.AddRow(color.MethodColor(r.Method), color.Yellow(r.Path), color.Cyan(r.Name))
	}
	table.Wrap = true

	// Print the table
	fmt.Print("\n ------------- Routes --------------- \n\n")
	fmt.Println(table)
	fmt.Print("\n")
}
