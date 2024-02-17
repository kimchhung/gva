package controller

import (
	"github.com/kimchhung/gva/app/module/admin/service"

	"github.com/kimchhung/gva/internal/control_route"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = (*AuthController)(nil)

type AuthController struct {
	service *service.AdminService
}

func (con *AuthController) Routes(r fiber.Router) {
	r.Route(
		"/admins/auth",

		func(router fiber.Router) {

		},
	)
}

func NewAuthController(service *service.AdminService) *AuthController {
	return &AuthController{
		service: service,
	}
}
