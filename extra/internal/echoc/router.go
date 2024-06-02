package echoc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	RouteMeta struct {
		method string
		path   string
		name   string

		middlewares []echo.MiddlewareFunc
		handlers    []echo.HandlerFunc
	}

	MetaHandler = func() []echo.HandlerFunc
	MetaFunc    = func(meta *RouteMeta) MetaHandler
)

// Set add middlewares to current route
func (r *RouteMeta) Use(middlewares ...echo.MiddlewareFunc) *RouteMeta {
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

func (r *RouteMeta) Do(more ...echo.HandlerFunc) MetaHandler {
	return func() []echo.HandlerFunc {
		return more
	}
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
func (r *RouteMeta) DoWithScope(handler MetaHandler) MetaHandler {
	return handler
}

// Register registers routes defined by the controller methods.
func Register(app *echo.Group, controller Controller) {

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
				name:   strings.Replace(fmt.Sprintf("%v.%s", controllerType, method.Name), "*", "", 1),
				method: "GET",
				path:   "/",
			}

			meta.handlers = metaHandlerFunc(meta)()

			// fiber endpoint
			AddRoute(r, meta)
		}
	}
}

func AddRoute(r *echo.Group, meta *RouteMeta) {
	handler := func(c echo.Context) error {
		for _, handler := range meta.handlers {
			if err := handler(c); err != nil {
				return err
			}
		}
		return nil
	}

	// fiber auto register route head with get
	// uncomment if this if need

	// if meta.method == http.MethodGet {
	// 	r.Get(meta.path, append(meta.middlewares, handler)...).Name(meta.name)
	// 	return
	// }

	e := r.Add(meta.method, meta.path, handler, meta.middlewares...)
	e.Name = meta.name
}
