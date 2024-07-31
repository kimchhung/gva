package routeutil

import (
	"fmt"
	"slices"

	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/color"

	"github.com/gosuri/uitable"
	"github.com/labstack/echo/v4"
)

func GroupRouteToNested(flatMenuList []*ent.Menu) (parentRoutes []*ent.Menu) {
	parentMapRoute := make(map[pxid.ID][]*ent.Menu)

	for _, route := range flatMenuList {
		if route.Pid != nil {
			parentMapRoute[*route.Pid] = append(parentMapRoute[*route.Pid], route)
		}
	}

	for _, menu := range flatMenuList {
		menu.Edges.Children = AppendRouteChildrens(menu.ID, parentMapRoute)

		if menu.Pid == nil {
			parentRoutes = append(parentRoutes, menu)
		}
	}

	slices.SortFunc(parentRoutes, func(a *ent.Menu, b *ent.Menu) int {
		return a.Order - b.Order
	})

	return parentRoutes
}

func FlattenNestedMenu(parentRoutes []*ent.Menu) (flatRoutes []*ent.Menu) {
	for _, p := range parentRoutes {
		flatRoutes = append(flatRoutes, p)

		if len(p.Edges.Children) > 0 {
			flatRoutes = append(flatRoutes, FlattenNestedMenu(p.Edges.Children)...)
		}
	}

	return flatRoutes
}

func AppendRouteChildrens(Pid pxid.ID, parentMapRoute map[pxid.ID][]*ent.Menu) (children []*ent.Menu) {
	children, ok := parentMapRoute[Pid]
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
