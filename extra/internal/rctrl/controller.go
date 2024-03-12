package rctrl

// route control

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/config"
)

type ModuleRouter interface {
	Register(app fiber.Router, cfg *config.Config)
}

/*
requirement usage of controller

	Init(r fiber.Router) fiber.Router{
		return r.group("")
		return r
	}

	FuncName(meta *rctrl.RouteMeta) rctrl.MetaHandler{
		return meta.Get("/path").Do(...)
		return meta.Get("/path").DoWithScope(...)
	}
*/
type Controller interface {
	Init(r fiber.Router) fiber.Router
}
