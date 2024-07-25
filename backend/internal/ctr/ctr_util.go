package ctr

import (
	"fmt"
	"reflect"
	"strings"
)

func ReflectController(controller any) (*Ctr, error) {
	var (
		controllerMeta *Ctr
	)

	if ctr, ok := controller.(CTR); ok {
		controllerMeta = ctr.Init()
	} else if ctrw, ok := controller.(CTRWith); ok {
		c := New() // Assuming New() creates a new Ctr instance
		controllerMeta = ctrw.Init(c)
	} else {
		return nil, fmt.Errorf("controller:%v is missing init method", reflect.TypeOf(controller).Elem())
	}

	controllerType := reflect.TypeOf(controller)
	controllerValue := reflect.ValueOf(controller)

	for i := controllerType.NumMethod() - 1; i >= 0; i-- {
		method := controllerType.Method(i)

		if method.Type.NumOut() == 1 && method.Type.Out(0).AssignableTo(reflect.TypeOf((*Route)(nil))) {
			route := &Route{
				name:   strings.Replace(fmt.Sprintf("%v.%s", controllerType, method.Name), "*", "", 1),
				method: "GET",
				path:   "/",
			}

			defineRoute, ok := controllerValue.MethodByName(method.Name).Interface().(func(*Route) *Route)
			if !ok {
				continue
			}
			route = defineRoute(route)
			controllerMeta.routes = append(controllerMeta.routes, route)
		}
	}

	return controllerMeta, nil
}

func GroupController(flats ...*Ctr) (nested []*Ctr) {
	rootMap := make(map[string]*Ctr)

	for _, node := range flats {
		rootMap[node.ID] = node
	}

	for _, node := range flats {
		parent, foundParent := rootMap[node.PID]
		if foundParent {
			parent.Children = append(parent.Children, node)
		} else {
			nested = append(nested, node)
		}
	}

	return nested
}

func FlatController(root ...*Ctr) ([]*Ctr, error) {
	visited := make(map[string]bool) // Tracks visited nodes to prevent infinite recursion
	var flats []*Ctr

	for _, r := range root {
		if err := flattenNode(r, &flats, visited); err != nil {
			return nil, err
		}
	}

	return flats, nil
}

func flattenNode(node *Ctr, flats *[]*Ctr, visited map[string]bool) error {
	if visited[node.ID] {
		return fmt.Errorf("circular reference detected for node %s", node.ID)
	}
	visited[node.ID] = true

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			if err := flattenNode(child, flats, visited); err != nil {
				return err
			}
		}
	}

	node.Children = []*Ctr{}
	node.parent = nil
	node.root = nil
	*flats = append(*flats, node)
	return nil
}
