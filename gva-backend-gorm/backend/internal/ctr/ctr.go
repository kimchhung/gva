package ctr

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type H = echo.HandlerFunc
type M = echo.MiddlewareFunc
type R = func(*echo.Group)

type ScopeHandler func() []H

type CTR interface {
	Init() *Ctr
}

type (
	Route struct {
		ctr    *Ctr
		method string
		path   string
		name   string

		Middlewares  []M
		ScopeHandler ScopeHandler
		Callback     R
	}
)

type RouteOption func(*Route) *Route

func NewRoute(opts ...RouteOption) *Route {
	r := &Route{
		method: http.MethodGet,
	}

	for _, o := range opts {
		r = o(r)
	}

	return r
}

func (r *Route) Use(middlewares ...M) *Route {
	r.Middlewares = append(r.Middlewares, middlewares...)
	return r
}

func (r *Route) Method(method string) *Route {
	r.method = method
	return r
}

func (r *Route) GetMethod() string {
	return r.method
}

func (r *Route) GetName() string {
	return r.name
}

func (r *Route) GetPath() string {
	if r.path == "" {
		return ""
	}

	return "/" + r.path
}

func (r *Route) GetFullPath() string {
	return r.ctr.GetPrefix() + r.GetPath()
}

func (r *Route) Path(path string) *Route {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	r.path = path
	return r
}

func (r *Route) Name(name string) *Route {
	r.name = name
	return r
}

func (r *Route) Do(scopeHandler func() []H) *Route {
	r.ScopeHandler = scopeHandler
	return r
}

func Add(callback R) *Route {
	r := NewRoute()
	r.Callback = callback
	return r
}

// GET
func GET(path string) *Route {
	r := NewRoute()
	return r.Method(http.MethodGet).Path(path)
}

// POST
func POST(path string) *Route {
	r := NewRoute()
	return r.Method(http.MethodPost).Path(path)
}

// PUT
func PUT(path string) *Route {
	r := NewRoute()
	return r.Method(http.MethodPut).Path(path)
}

// DELETE
func DELETE(path string) *Route {
	r := NewRoute()
	return r.Method(http.MethodDelete).Path(path)
}

// OPTIONS
func OPTIONS(path string) *Route {
	r := NewRoute()
	return r.Method(http.MethodOptions).Path(path)
}

// PATCH
func PATCH(path string) *Route {
	r := NewRoute()
	return r.Method(http.MethodPatch).Path(path)
}

type Middleware interface {
	Name() string
}

type Ctr struct {
	prefix      string
	Routes      []*Route
	Middlewares []M
}

type CtrOption func(*Ctr) *Ctr

func New(opts ...CtrOption) *Ctr {
	c := &Ctr{}

	for _, o := range opts {
		c = o(c)
	}

	return c
}

func Group(prefix string, middleware ...M) CtrOption {
	return func(c *Ctr) *Ctr {
		c.Prefix(prefix)
		c.Middlewares = append(c.Middlewares, middleware...)
		return c
	}
}

func Use(middleware ...M) CtrOption {
	return func(c *Ctr) *Ctr {
		c.Middlewares = append(c.Middlewares, middleware...)
		return c
	}
}

func (c *Ctr) IsEmptyPrefix() bool {
	switch c.prefix {
	case "", "/":
		return true
	}
	return false
}

func (c *Ctr) GetPrefix() string {
	if c.prefix == "" {
		return c.prefix
	}

	return "/" + c.prefix
}

func (c *Ctr) Prefix(prefix string) {
	prefix = strings.TrimPrefix(prefix, "/")
	prefix = strings.TrimSuffix(prefix, "/")
	c.prefix = prefix
}
