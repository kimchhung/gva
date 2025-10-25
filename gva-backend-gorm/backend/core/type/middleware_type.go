package coretype

import "github.com/labstack/echo/v4"

type MiddlewareRouter interface {
	Use(middleware ...echo.MiddlewareFunc)
}

type Middleware interface {
	RegisterMiddleware(c MiddlewareRouter)
}
