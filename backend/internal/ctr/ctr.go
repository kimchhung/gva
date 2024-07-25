package ctr

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gva/internal/treeprint"
	"github.com/labstack/echo/v4"
)

type H = echo.HandlerFunc
type M = echo.MiddlewareFunc

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
	group    string
	routes   []*Route
	parent   *Ctr
	children []*Ctr

	groupOption []func(*echo.Group) *echo.Group
}

func New() *Ctr {
	return &Ctr{}
}

func (c *Ctr) ForEach(f func(c *Ctr)) *Ctr {
	f(c)

	for _, child := range c.children {
		child.ForEach(f)
	}

	return c
}

func (c *Ctr) PrintTree(prefix string) {

	root := printTree(c, nil)
	treeprint.Print(root)
}

func (c *Ctr) SetGroup(g string) *Ctr {
	c.group = strings.ReplaceAll(g, "/", "")
	return c
}

func (c *Ctr) Group() string {
	return c.group
}

func (c *Ctr) FullPath() string {
	return getGroup(c)
}

func (c *Ctr) Add(controller CTR) {
	// Reflectively analyze the controller and generate its metadata
	controllerMeta := c.reflectController(controller, nil)
	c.children = append(c.children, controllerMeta...)
}

func (c *Ctr) reflectController(controller CTR, parent *Ctr) (list []*Ctr) {
	controllerMeta := &Ctr{}
	controllerMeta = controller.Init(controllerMeta)
	if parent != nil {
		controllerMeta.parent = parent
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

	// Handle children recursively if the controller implements Children
	if childrenController, ok := controller.(Children); ok {
		for _, child := range childrenController.Children() {
			childMeta := c.reflectController(child, controllerMeta) // Set current controllerMeta as parent for child
			controllerMeta.children = append(controllerMeta.children, childMeta...)
		}
	}

	list = append(list, controllerMeta)
	return list
}

type CTR interface {
	Init(ctm *Ctr) *Ctr
}

type Parent interface {
	Parent() CTR
}

type Children interface {
	Children() []CTR
}

func getGroup(c *Ctr) (path string) {
	if c.parent == nil {
		path += c.group
	} else {
		path += getGroup(c.parent) + "/" + c.group
	}

	return
}

func printTree(c *Ctr, parentNode *treeprint.Node) *treeprint.Node {
	if parentNode == nil {
		parentNode = treeprint.New(c.group)
	} else {
		parentNode = parentNode.Add(fmt.Sprintf("%v", c.group))
	}

	for _, r := range c.routes {
		parentNode.Add(fmt.Sprintf("%s %s", r.method, r.path))
	}

	for _, ch := range c.children {
		printTree(ch, parentNode)
	}

	return parentNode
}
