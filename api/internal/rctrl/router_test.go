package rctrl_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/internal/rctrl"
)

// BenchmarkRegisterRoutes measures the performance of the Register function.

// Define a simple controller with a single method that returns a MetaHandler

type MyController struct{}

func (con *MyController) Hello(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/hello").Name("get many roles").Do(func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, World! %v", con))
	})
}

func BenchmarkRegisterRoutes(b *testing.B) {
	app := fiber.New()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		rctrl.Register(app, &MyController{})
	}
}

// BenchmarkHandleRequestCtr-16    	   39570	     30687 ns/op	    5843 B/op	      27 allocs/op
func BenchmarkHandleRequestCtr(b *testing.B) {
	app := fiber.New()
	rctrl.Register(app, &MyController{})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}

// BenchmarkHandleRequest-16    	   30840	     33726 ns/op	    5828 B/op	      25 allocs/op
func BenchmarkHandleRequest(b *testing.B) {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, World! %v", &MyController{}))
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}
