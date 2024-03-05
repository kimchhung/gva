package module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/rctrl"
)

type Router struct {
	controllers []rctrl.Controller
}

func NewRouter(controllers ...rctrl.Controller) *Router {
	return &Router{
		controllers,
	}
}

func (r *Router) Register(app fiber.Router) {
	for _, controller := range r.controllers {
		controller.Routes(app)
	}
}
