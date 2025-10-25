package swagger

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Register(app *echo.Echo, basePath string, path string, config ...func(*echoSwagger.Config)) {
	swaggerPath := basePath + path
	swaggerRedirect := func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, swaggerPath+"/index.html")
	}

	app.GET(swaggerPath, swaggerRedirect)
	app.GET(swaggerPath+"/", swaggerRedirect)
	app.GET(swaggerPath+"/*", echoSwagger.EchoWrapHandler(config...))
}
