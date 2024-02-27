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
	auth := r.Group("auth")
	rctrl.Register(auth, con)
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
// @Router /auth/login [post]
func (con *AuthController) Login(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/login").Name("admin login").DoWithScope(func() []fiber.Handler {
		body := new(dto.LoginRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				token, admin, err := con.service.LoginAdmin(c.UserContext(), body.Username, body.Password)
				if err != nil {
					return err
				}

				res := dto.LoginResponse{
					Token: token,
					Admin: admin,
				}

				return request.Resp(c, request.Response{
					Message: "The admin was logined successfully!",
					Data:    res,
				})
			},
		}
	})
}

// @Tags Auth
// @Summary Register a new admin
// @Description Register a new admin with the provided credentials
// @ID create-Auth-register
// @Accept  json
// @Produce  json
// @Param Auth body dto.RegisterRequest true "Registration data"
// @Success  200 {object} request.Response{data=dto.RegisterResponse} "Successfully registered admin"
// @Router /auth/register [post]
func (con *AuthController) Register(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/register").Name("admin register").DoWithScope(func() []fiber.Handler {
		body := new(dto.RegisterRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				// Assuming RegisterAdmin returns a user object and an error
				token, admin, err := con.service.RegisterAdmin(c.UserContext(), body.Username, body.Password, body.DisplayName)
				if err != nil {
					return err
				}

				data := dto.RegisterResponse{
					Token: token,
					Admin: admin,
				}

				// Assuming you want to return some user data in the response
				return request.Resp(c, request.Response{
					Message: "The admin was registered successfully!",
					Data:    data, // Adjust this based on what you want to return
				})
			},
		}
	})
}
