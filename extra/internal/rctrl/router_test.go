package rctrl_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/rctrl"
)

// BenchmarkRegisterRoutes measures the performance of the Register function.

// Define a simple controller with a single method that returns a MetaHandler

type MyController struct{}

func (MyController) Init(r fiber.Router) fiber.Router {
	return r
}

// 743730	      1660 ns/op	    1465 B/op	      30 allocs/op
func (con *MyController) Hello(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/hello").Do(func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

// BenchmarkRegisterRoutes-8   	  634002	      1708 ns/op	    1499 B/op	      30 allocs/op
func BenchmarkRegisterRoutes(b *testing.B) {
	app := fiber.New()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		rctrl.Register(app, &MyController{})
	}
}

// BenchmarkHandleRequestCtr-8   	  220992	      5797 ns/op	    5773 B/op	      25 allocs/op
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

// BenchmarkHandleRequest-8   	  229569	      5311 ns/op	    5776 B/op	      25 allocs/op
func BenchmarkHandleRequest(b *testing.B) {
	app := fiber.New()

	g := app.Group("/hello")
	g.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}
