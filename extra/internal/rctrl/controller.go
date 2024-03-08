package rctrl

// route control

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/config"
)

type Controller interface {
	Routes(router fiber.Router)
}

type ModuleRouter interface {
	Register(app fiber.Router, cfg *config.Config)

	// dashbaord | web
	Name() string
}
