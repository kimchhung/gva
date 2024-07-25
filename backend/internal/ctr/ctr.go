package ctr

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type H = echo.HandlerFunc
type M = func(next H) error

type ScopeHandler func() []H

type (
	Route struct {
		method string
		path   string
		name   string

		middlewares  []M
		handlers     []H
		scopeHandler ScopeHandler
	}
)

// Set add middlewares to current route
func (r *Route) Use(middlewares ...M) *Route {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

// Set Method to current route
func (r *Route) Method(method string) *Route {
	r.method = method
	return r
}

// Set path to current route
func (r *Route) Path(path string) *Route {
	r.path = path
	return r
}

func (r *Route) Name(name string) *Route {
	r.name = name
	return r
}

// MethodPost sets the method to POST and the path for the route.
func (r *Route) Get(path string) *Route {
	return r.Method(http.MethodGet).Path(path)
}

// MethodPost sets the method to POST and the path for the route.
func (r *Route) Post(path string) *Route {
	return r.Method(http.MethodPost).Path(path)
}

// MethodPut sets the method to PUT and the path for the route.
func (r *Route) Put(path string) *Route {
	return r.Method(http.MethodPut).Path(path)
}

// MethodDelete sets the method to DELETE and the path for the route.
func (r *Route) Delete(path string) *Route {
	return r.Method(http.MethodDelete).Path(path)
}

// MethodOptions sets the method to OPTIONS and the path for the route.
func (r *Route) Options(path string) *Route {
	return r.Method(http.MethodOptions).Path(path)
}

// MethodPatch sets the method to PATCH and the path for the route.
func (r *Route) Patch(path string) *Route {
	return r.Method(http.MethodPatch).Path(path)
}

func (r *Route) Pre(hnadlers ...H) *Route {
	r.handlers = append(r.handlers, hnadlers...)
	return r
}

func (r *Route) Do(more ...H) *Route {
	r.Pre(more...)
	r.scopeHandler = func() []H {
		return r.handlers
	}
	return r
}

func (r *Route) DoWithScope(handler ScopeHandler) *Route {
	r.scopeHandler = handler
	return r
}

func (r *Route) ScopeHandler(handler ScopeHandler) ScopeHandler {
	return r.scopeHandler
}

/*
provide a scope to store middleware data and reuse

	meta.Path("/all").MethodPost()
	return meta.DoWithScope(func() []echo.HandlerFunc {
		// * middlewares safe storage
		var ip string
		return []echo.HandlerFunc{
			ParseIp(&ip), // parsers

			func(c echo.Context) error { // handler
				return c.SendString(ip)
			},
		}
	})
*/

type Ctr struct {
	ID       string `json:"id"`
	PID      string `json:"pid"`
	Children []*Ctr
	parent   *Ctr
	root     *Ctr

	prefix  string
	groupFn func(*echo.Group) *echo.Group
	routes  []*Route
}

func (c *Ctr) IsRoot() bool {
	return c.PID != ""
}

func (c *Ctr) Root() *Ctr {
	return c.root
}

func New(opts ...Option) *Ctr {
	c := &Ctr{}
	c.Option(opts...)
	return c
}

func (c *Ctr) Option(opts ...Option) *Ctr {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Option func(*Ctr)

func Group(prefix string, middleware ...echo.MiddlewareFunc) Option {
	return func(c *Ctr) {
		c.prefix = prefix
		c.groupFn = func(g *echo.Group) *echo.Group {
			return g.Group(prefix, middleware...)
		}
	}
}

func ParentID(parentId any) Option {
	return func(c *Ctr) {
		c.PID = fmt.Sprintf("%v", parentId)
	}
}

func (c *Ctr) Add(controller any) (err error) {
	ctr, ok := controller.(*Ctr)
	if ok {
		if c.IsRoot() {
			ctr.root = c
		} else {
			ctr.root = c.Root()
		}

		ctr.PID = c.ID
		ctr.parent = c
		ctr.Children = append(ctr.Children, ctr)
		return nil
	}

	// Reflectively analyze the controller and generate its metadata
	ctr, err = ReflectController(controller)
	if err != nil {
		if c.IsRoot() {
			ctr.root = c
		} else {
			ctr.root = c.Root()
		}

		ctr.PID = c.ID
		ctr.parent = c
		c.Children = append(c.Children, ctr)
		return err
	}
	return nil
}

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
		parent, ok := rootMap[node.PID]
		if ok {
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
	*flats = append(*flats, node)
	return nil
}

type CTRWith interface {
	Init(ctm *Ctr) *Ctr
}

type CTR interface {
	Init() *Ctr
}
