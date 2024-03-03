package request_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/utils/request"
)

/*
goos: darwin
goarch: arm64
pkg: github.com/kimchhung/gva/extra/utils/request
BenchmarkRequestData-8   	  129326	      9023 ns/op	    6439 B/op	      35 allocs/op
*/
func BenchmarkRequestData(b *testing.B) {
	app := fiber.New()

	list := []string{"123", "123", "123"}
	app.Get("/hello", func(c *fiber.Ctx) error {
		return request.Resp(c,
			request.Data(list),
		)
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}

/*
goos: darwin
goarch: arm64
pkg: github.com/kimchhung/gva/extra/utils/request
BenchmarkRequestListData-8   	  134776	      8606 ns/op	    6438 B/op	      35 allocs/op
*/
func BenchmarkRequestListData(b *testing.B) {
	app := fiber.New()

	list := []string{"123", "123", "123"}
	app.Get("/hello", func(c *fiber.Ctx) error {
		return request.Resp(c,
			request.Data(list),
		)
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}

type TestStruct struct {
	Me  string `json:"me"`
	You string `json:"you"`
}

/*
goos: darwin
goarch: arm64
pkg: github.com/kimchhung/gva/extra/utils/request
BenchmarkRequestStruct-8   	  137684	      7906 ns/op	    5973 B/op	      28 allocs/op
PASS
ok  	github.com/kimchhung/gva/extra/utils/request	2.792s
*/
func BenchmarkRequestStruct(b *testing.B) {
	app := fiber.New()

	TestStructData := &TestStruct{
		Me:  "hahhaa",
		You: "hahaha",
	}

	app.Get("/hello", func(c *fiber.Ctx) error {
		return request.Resp(c,
			request.Data(TestStructData),
		)
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}

/*
goos: darwin
goarch: arm64
pkg: github.com/kimchhung/gva/extra/utils/request
BenchmarkRequestMap-8   	  147314	      9363 ns/op	    6607 B/op	      38 allocs/op
PASS
*/
func BenchmarkRequestMap(b *testing.B) {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return request.Resp(c,
			request.Data(map[string]any{
				"Me":  "hahhaa",
				"You": "hahaha",
			}),
		)
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}

/*
goos: darwin
goarch: arm64
pkg: github.com/kimchhung/gva/extra/utils/request
BenchmarkRequestMapOriginal-8   	  121930	      8482 ns/op	    6592 B/op	      37 allocs/op
*/
func BenchmarkRequestMapOriginal(b *testing.B) {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(
			request.Response{
				Code:    200,
				Message: "ok",
				Data: map[string]any{
					"Me":  "hahhaa",
					"You": "hahaha",
				},
				HttpStatus: 200,
			},
		)
	})

	// Create a new HTTP request with the route from the test case
	req := httptest.NewRequest("GET", "/hello", nil)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req)
	}
}
