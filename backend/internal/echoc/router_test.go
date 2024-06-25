package echoc_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gva/internal/echoc"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// BenchmarkRegisterRoutes measures the performance of the Register function.

// Define a simple controller with a single method that returns a MetaHandler

var _ interface{ echoc.Controller } = (*MyController)(nil)

type MyController struct{}

func (MyController) Init(r *echo.Group) *echo.Group {
	return r
}

func (con *MyController) Hello(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/hello").Do(func(c echo.Context) error {
		return c.String(200, "hello")
	})
}

// BenchmarkRegisterRoutes-8   	  634002	      1708 ns/op	    1499 B/op	      30 allocs/op
func BenchmarkRegisterRoutes(b *testing.B) {
	app := echo.New().Group("/")

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		echoc.Register(app, &MyController{})
	}
}

// BenchmarkHandleRequestCtr-16    	  137224	     10644 ns/op	    6916 B/op	      31 allocs/op
func BenchmarkHandleRequestCtr(b *testing.B) {
	e := echo.New()
	app := e.Group("/")

	con := &MyController{}
	echoc.Register(app, con)

	helloHandlers := con.Hello(&echoc.RouteMeta{})
	handler := func(cc echo.Context) (err error) {
		for _, h := range helloHandlers() {
			if err = h(cc); err == nil {
				return err
			}
		}
		return
	}

	// Create a new HTTP request with the route from the test case

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/hello", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/hello")

		if err := handler(c); err != nil {
			b.Error(err)
		}

		res := rec.Result()
		assert.Equal(b, http.StatusOK, res.StatusCode, "Expected status code 200")

		// Correctly read and convert the response body to a string
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			b.Fatal(err)
		}
		bodyString := string(bodyBytes)
		assert.Equal(b, "hello", bodyString, "Expected response body to be 'hello'")

		defer res.Body.Close()
	}
}
