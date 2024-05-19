package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/api/admin/module/auth/dto"
	"github.com/kimchhung/gva/extra/app/common/service"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

var _ interface{ rctrl.Controller } = (*AuthController)(nil)

type AuthController struct {
	service *AuthService
	jwt_s   *service.JwtService
}

func (con *AuthController) Init(r fiber.Router) fiber.Router {
	return r.Group("auth")
}

func NewAuthController(service *AuthService, jwtService *service.JwtService) *AuthController {
	return &AuthController{
		service: service,
		jwt_s:   jwtService,
	}
}

// @Tags			Auth
// @Summary		Authenticate a admin
// @Description	Authenticate a admin with the provided credentials
// @ID				create-Auth
// @Accept			json
// @Produce		json
// @Param			Auth	body		dto.LoginRequest																true	"Auth data"
// @Success		200		{object}	response.Response{data=map[string]dto.LoginResponse{list=[]dto.LoginResponse}}	"Successfully created Auth"
// @Router			/auth/login [post]
func (con *AuthController) Login(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/login").DoWithScope(func() []fiber.Handler {
		body := new(dto.LoginRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				token, admin, err := con.service.LoginAdmin(c.UserContext(), body)
				if err != nil {
					return err
				}

				res := dto.LoginResponse{
					Token: token,
					Admin: admin,
				}

				return request.Response(c,
					response.Data(res),
					response.Message("The admin was logined successfully!"),
				)
			},
		}
	})
}

// @Tags			Auth
// @Summary		Register a new admin
// @Description	Register a new admin with the provided credentials
// @ID				create-Auth-register
// @Accept			json
// @Produce		json
// @Param			Auth	body		dto.RegisterRequest								true	"Registration data"
// @Success		200		{object}	response.Response{data=dto.RegisterResponse}	"Successfully registered admin"
// @Router			/auth/register [post]
func (con *AuthController) Register(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/register").DoWithScope(func() []fiber.Handler {
		body := new(dto.RegisterRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				// Assuming RegisterAdmin returns a user object and an error
				token, admin, err := con.service.RegisterAdmin(c.UserContext(), body)
				if err != nil {
					return err
				}

				data := dto.RegisterResponse{
					Token: token,
					Admin: admin,
				}

				// Assuming you want to return some user data in the response
				return request.Response(c,
					response.Data(data),
					response.Message("The admin was registered successfully!"),
				)
			},
		}
	})
}

// @Tags		Auth
// @Security	Bearer
// @Summary		query your self after login
// @Description	get admin data by token
// @ID			create-Auth-me
// @Accept		json
// @Produce		json
// @Success		200	{object}	response.Response{data=ent.Admin}	"Successfully registered admin"
// @Router			/auth/me [get]
func (con *AuthController) Me(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	meta.Use(
		con.jwt_s.RequiredAdmin(),
	)
	return meta.Get("/me").Use(con.jwt_s.RequiredAdmin()).DoWithScope(func() []fiber.Handler {
		admin := new(ent.Admin)

		return []fiber.Handler{
			// extract the admin from context which inject by jwt_s.RequiredAdmin()
			request.MustAdmin(admin),

			// handler
			func(c *fiber.Ctx) error {
				return request.Response(c, response.Data(admin))
			},
		}
	})
}
