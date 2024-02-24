package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/module/auth/dto"
	"github.com/kimchhung/gva/app/module/auth/service"
	"github.com/kimchhung/gva/internal/rctrl"
	"github.com/kimchhung/gva/utils/request"
)

// don't remove for runtime type checking
var _ IAuthController = (*AuthController)(nil)

type IAuthController interface {
	rctrl.FiberRouter
	Login(meta *rctrl.RouteMeta) rctrl.MetaHandler
}

type AuthController struct {
	service *service.AuthService
}

func (con *AuthController) Routes(r fiber.Router) {
	r.Route("auth",
		func(router fiber.Router) {
			rctrl.Register(router, con)
		},
	)
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// @Tags Auth
// @Summary Authenticate a admin
// @Description Authenticate a admin with the provided credentials
// @ID create-Auth
// @Accept  json
// @Produce  json
// @Param Auth body dto.LoginRequest true "Auth data"
// @Success  200 {object} request.Response{data=dto.LoginResponse} "Successfully created Auth"
// @Router /auth [post]
func (con *AuthController) Login(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/admin-login").Name("admin login").DoWithScope(func() []fiber.Handler {
		body := new(dto.LoginRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {

				return request.Resp(c, request.Response{
					Message: "The auth was created successfully!",
					Data:    "",
				})
			},
		}
	})
}
