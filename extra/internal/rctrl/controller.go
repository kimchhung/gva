package rctrl

// route control

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Routes(router fiber.Router)
}
