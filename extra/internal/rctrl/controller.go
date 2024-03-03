package rctrl

// route control

import "github.com/gofiber/fiber/v2"

type FiberRouter interface {
	Routes(router fiber.Router)
}

type Router interface {
	Register()
}
