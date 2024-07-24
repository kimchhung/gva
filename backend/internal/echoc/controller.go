package echoc

// route control

import (
	"github.com/labstack/echo/v4"
)

type ModuleRouter interface {
	Register(app *echo.Echo, args ...any)
}

type Controller interface {
	Init(r *echo.Group) *echo.Group
}

type ControllerParent interface {
	Parent() Controller
}

type Group struct {
	*echo.Group
}
