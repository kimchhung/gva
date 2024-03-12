package rctrl

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type (
	RouteMeta struct {
		method string
		path   string
		name   string

		middlewares []fiber.Handler
		handlers    []fiber.Handler
	}

	MetaHandler = func() []fiber.Handler
	MetaFunc    = func(meta *RouteMeta) MetaHandler
)

// Set add middlewares to current route
func (r *RouteMeta) Use(middlewares ...fiber.Handler) *RouteMeta {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

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
			ParseIp(&ip), // parsers

			func(c *fiber.Ctx) error { // handler
				return c.SendString(ip)
			},
		}
	})
*/
func (r *RouteMeta) DoWithScope(handler MetaHandler) MetaHandler {
	return handler
}

// Register registers routes defined by the controller methods.
func Register(app fiber.Router, controller Controller) {
	r := controller.Init(app)
	controllerType := reflect.TypeOf(controller)
	controllerValue := reflect.ValueOf(controller)

	for i := controllerType.NumMethod() - 1; i >= 0; i-- {
		method := controllerType.Method(i)
		if method.Type.NumOut() == 1 && method.Type.Out(0).ConvertibleTo(reflect.TypeOf((*MetaHandler)(nil)).Elem()) {
			methodValue := controllerValue.MethodByName(method.Name)
			metaHandlerFunc, ok := methodValue.Interface().(func(*RouteMeta) MetaHandler)
			if !ok {
				// Log an error instead of panicking
				log.Printf("controller method %s must be a func(*RouteMeta) MetaHandler", method.Name)
				continue
			}

			meta := &RouteMeta{
				name:   fmt.Sprintf("%s.%s", controllerType.Elem().Name(), method.Name),
				method: "GET",
				path:   "/",
			}

			meta.handlers = metaHandlerFunc(meta)()
			AddRoute(r, meta)
		}
	}
}

func AddRoute(r fiber.Router, meta *RouteMeta) {
	handler := func(c *fiber.Ctx) error {
		for _, handler := range meta.handlers {
			if err := handler(c); err != nil {
				return err
			}
		}
		return nil
	}

	if meta.method == http.MethodGet {
		r.Get(meta.path, append(meta.middlewares, handler)...).Name(meta.name)
		return
	}

	r.Add(meta.method, meta.path, append(meta.middlewares, handler)...).Name(meta.name)
}
