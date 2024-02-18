package rctrl

import (
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type (
	RouteMeta struct {
		method string
		path   string
		name   string
	}

	MetaHandler = func() []fiber.Handler
	MetaFunc    = func(meta *RouteMeta) MetaHandler
)

// Set Method to current route
func (r *RouteMeta) Method(method string) *RouteMeta {
	r.method = method
	return r
}

// Set path to current route
func (r *RouteMeta) Path(path string) *RouteMeta {
	r.path = path
	return r
}

func (r *RouteMeta) Name(name string) *RouteMeta {
	r.name = name
	return r
}

// MethodPost sets the method to POST and the path for the route.
func (r *RouteMeta) Get(path string) *RouteMeta {
	return r.Method(http.MethodGet).Path(path)
}

// MethodPost sets the method to POST and the path for the route.
func (r *RouteMeta) Post(path string) *RouteMeta {
	return r.Method(http.MethodPost).Path(path)
}

// MethodPut sets the method to PUT and the path for the route.
func (r *RouteMeta) Put(path string) *RouteMeta {
	return r.Method(http.MethodPut).Path(path)
}

// MethodDelete sets the method to DELETE and the path for the route.
func (r *RouteMeta) Delete(path string) *RouteMeta {
	return r.Method(http.MethodDelete).Path(path)
}

// MethodOptions sets the method to OPTIONS and the path for the route.
func (r *RouteMeta) Options(path string) *RouteMeta {
	return r.Method(http.MethodOptions).Path(path)
}

// MethodPatch sets the method to PATCH and the path for the route.
func (r *RouteMeta) Patch(path string) *RouteMeta {
	return r.Method(http.MethodPatch).Path(path)
}

func (r *RouteMeta) Do(fn fiber.Handler, more ...fiber.Handler) MetaHandler {
	return func() []func(*fiber.Ctx) error {
		return append([]fiber.Handler{fn}, more...)
	}
}

/*
provide a scope to store middleware data and reuse

	meta.Path("/all").MethodPost()
	return meta.DoWithScope(func() []fiber.Handler {
		// * middlewares safe storage
		var ip string
		return []fiber.Handler{
			ParseIp(&ip), // middlewares

			func(c *fiber.Ctx) error { // handler
				return c.SendString(ip)
			},
		}
	})
*/
func (r *RouteMeta) DoWithScope(handler MetaHandler) MetaHandler {
	return handler
}

func defineRoute(app fiber.Router, r *RouteMeta, defineMeta func(meta *RouteMeta) MetaHandler) fiber.Router {
	getHandlers := defineMeta(r)
	fn := func(c *fiber.Ctx) error {
		var err error
		for _, handler := range getHandlers() {
			if err = handler(c); err != nil {
				return err
			}
		}
		return err
	}

	if r.method == http.MethodGet {
		return app.Get(r.path, fn).Name(r.name)
	}

	return app.Add(r.method, r.path, fn).Name(r.name)
}

// default path is "/" and method is "GET"
func NewRoute(app fiber.Router, metaFunc MetaFunc) fiber.Router {
	return defineRoute(app, &RouteMeta{}, metaFunc)
}

// Register registers routes defined by the controller methods.
func Register(app fiber.Router, controller any) {
	controllerType := reflect.TypeOf(controller)

	// Iterate over the methods of the controller
	for i := 0; i < controllerType.NumMethod(); i++ {
		method := controllerType.Method(i)

		if method.Type.NumOut() == 1 && method.Type.Out(0).ConvertibleTo(reflect.TypeOf((*MetaHandler)(nil)).Elem()) {
			methodValue := reflect.ValueOf(controller).MethodByName(method.Name)
			metaHandler, ok := methodValue.Interface().(func(*RouteMeta) MetaHandler)
			if !ok {
				panic("controller method must be a func(*RouteMeta) MetaHandler")
			}

			// Create a new route using the MetaHandler
			NewRoute(app, metaHandler)
		}
	}
}
