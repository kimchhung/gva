package rctrl_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/rctrl"
)

// BenchmarkRegisterRoutes measures the performance of the Register function.

// Define a simple controller with a single method that returns a MetaHandler

type MyController struct{}

// 743730	      1660 ns/op	    1465 B/op	      30 allocs/op
func (con *MyController) Hello(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/hello").Do(func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

// BenchmarkRegisterRoutes-8   	 1921436	       603.7 ns/op	     981 B/op	      15 allocs/op
// BenchmarkRegisterRoutes-8   	  712566	      1614 ns/op	    1469 B/op	      30 allocs/op
func BenchmarkRegisterRoutes(b *testing.B) {
	app := fiber.New()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		rctrl.Register(app, &MyController{})
	}
}

// BenchmarkHandleRequestCtr-8   	  228070	      5311 ns/op	    5779 B/op	      25 allocs/op
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
	app := fiber.New(
		fiber.Config{
			EnablePrintRoutes: true,
		},
	)

	app.Get("/paginate", func(c *fiber.Ctx) error {
		fmt.Println("++++++")
		return c.SendString("Hello, World!")
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		fmt.Println("-----")
		return c.SendString("Hello, World!" + c.Params("id", ""))
	})

	//print ------

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/paginate", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}
