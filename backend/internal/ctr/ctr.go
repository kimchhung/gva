package ctr

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gva/internal/treeprint"
	"github.com/gva/utils/color"
	"github.com/labstack/echo/v4"
)

type H = echo.HandlerFunc
type M = func(next H) error

type ScopeHandler func() []H

type CTRWith interface {
	Init(ctm *Ctr) *Ctr
}

type CTR interface {
	Init() *Ctr
}

type (
	Route struct {
		method string
		path   string
		name   string

		middlewares  []M
		handlers     []H
		scopeHandler ScopeHandler
		ctr          *Ctr
	}
)

func NewRoute() *Route {
	return &Route{}
}

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

func (r *Route) FullPath() string {
	return r.ctr.FullPrefix() + r.path
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
	ID  string
	PID string

	Children []*Ctr
	parent   *Ctr
	root     *Ctr
	all      map[string]*Ctr

	prefix  string
	groupFn func(*echo.Group) *echo.Group
	routes  []*Route

	_infoFn      func(*Ctr) any
	_infoRouteFn func(*Route) any
}

func (c *Ctr) infoFn(ctr *Ctr) any {
	if c.root != nil {
		return c._infoFn(ctr)
	}
	return c._infoFn(ctr)
}

func (c *Ctr) infoRouteFn(r *Route) any {
	if c.root != nil {
		return c._infoRouteFn(r)
	}
	return c._infoRouteFn(r)
}

// check has in root
func (c *Ctr) Has(id any) bool {
	_, ok := c.GetAll()[fmt.Sprintf("%v", id)]
	return ok
}

// get all from root chain
func (c *Ctr) GetAll() map[string]*Ctr {
	if c.root != nil {
		if c.root.all == nil {
			c.root.all = map[string]*Ctr{}
		}
		return c.root.all
	}

	if c.all == nil {
		c.all = map[string]*Ctr{}
	}

	return c.all
}

// get one from root
func (c *Ctr) Get(id any) (*Ctr, bool) {
	d, ok := c.GetAll()[fmt.Sprintf("%v", id)]
	return d, ok
}

func (c *Ctr) set(ctr *Ctr) {
	c.GetAll()[ctr.ID] = ctr
}

// concurrent add not support
func (c *Ctr) AddChild(ctr *Ctr) (child *Ctr, err error) {
	if ctr.ID == "" {
		return nil, errors.New("id is required, use ctr.ID(any)")
	}

	if c.Has(ctr.ID) {
		return nil, fmt.Errorf("duplicated Id %v", ctr.ID)
	}
	c.set(ctr)

	ctr.PID = c.ID
	ctr.parent = c

	// Update root if necessary
	if c.parent == nil {
		c.root = c
	}

	ctr.root = c.root
	c.Children = append(c.Children, ctr)
	return ctr, nil
}

// concurrent add not support
func (c *Ctr) AddChildren(ctrs ...*Ctr) ([]*Ctr, error) {
	added := make([]*Ctr, len(ctrs))
	var err error

	for i, ctr := range ctrs {
		added[i], err = c.AddChild(ctr)
		if err != nil {
			return nil, err
		}
	}

	return added, nil
}

func (c *Ctr) Parent() *Ctr {
	return c.parent
}

func (c *Ctr) FullPrefix() string {
	return FullPathCtr(c)
}

func (c *Ctr) Root() *Ctr {
	return c.root
}

func DefaultInfoFn(c *Ctr) any {
	return color.Cyan(c.ID)
}

func DefaultRouteFn(r *Route) any {
	return fmt.Sprintf("%v %v", color.MethodColor(r.method), r.FullPath())
}

func New(opts ...Option) *Ctr {
	c := &Ctr{
		_infoFn:      DefaultInfoFn,
		_infoRouteFn: DefaultRouteFn,
	}
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
		c.ID = prefix
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

func AddRoute(routes ...*Route) Option {
	return func(c *Ctr) {
		for _, r := range routes {
			r.ctr = c
			c.routes = append(c.routes, r)
		}
	}
}

func ID(parentId any) Option {
	return func(c *Ctr) {
		c.ID = fmt.Sprintf("%v", parentId)
		c.prefix = c.ID
	}
}

func Each(c *Ctr, f func(c *Ctr)) {
	for _, child := range c.Children {
		Each(child, f)
	}
}

func CreateIDTreePrint(parent *treeprint.Node, list ...*Ctr) (current *treeprint.Node) {
	for _, c := range list {
		if parent == nil {
			current = treeprint.New(c.infoFn(c))
		} else {
			current = parent.Add(c.infoFn(c))
			for _, r := range c.routes {
				current.Add(c.infoRouteFn(r))
			}
		}

		for _, child := range c.Children {
			CreateIDTreePrint(current, child) // Corrected recursive call
		}
	}

	return current // Return the current node
}

func AddToParentNested(parent *Ctr, child *Ctr) (*Ctr, error) {
	childNode, err := parent.AddChild(child)
	if err != nil {
		return nil, err // Return the error immediately if adding a child fails
	}

	for _, node := range childNode.Children {
		childNode.AddChild(node)
	}

	return childNode, nil
}

func FullPathCtr(c *Ctr) (path string) {
	if c.parent == nil {
		return c.prefix
	}

	fullPath := FullPathCtr(c.parent) + c.prefix
	return fullPath
}
