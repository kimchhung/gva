package auth

import (
	"github.com/gva/api/admin/module/auth/dto"
	"github.com/gva/app/common/service"
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*AuthController)(nil)

type AuthController struct {
	service *AuthService
	jwt_s   *service.JwtService
}

func NewAuthController(service *AuthService, jwt_s *service.JwtService) *AuthController {
	return &AuthController{
		service: service,
		jwt_s:   jwt_s,
	}
}

func (con *AuthController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/auth", con.jwt_s.RequiredAdmin()),
	)
}

// @Tags			Auth
// @Summary			Authenticate a admin
// @Description		Authenticate a admin with the provided credentials
// @ID				create-Auth
// @Accept			json
// @Produce			json
// @Param			Auth	body		dto.LoginRequest	true	"Auth data"
// @Success			200		{object}	response.Response{data=dto.LoginResponse}	"Successfully created Auth"
// @Router			/auth/login [post]
func (con *AuthController) Login() *ctr.Route {
	return ctr.POST("/login").Do(func() []ctr.H {
		body := new(dto.LoginRequest)

		return []ctr.H{
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.LoginAdmin(c.Request().Context(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The admin was logined successfully!"),
				)
			},
		}
	})
}

// @Tags			Auth
// @Summary			Register a new admin
// @Description		Register a new admin with the provided credentials
// @ID				create-Auth-register
// @Accept			json
// @Produce			json
// @Param			Auth	body		dto.RegisterRequest		true	"Registration data"
// @Success			200		{object}	response.Response{data=dto.RegisterResponse}	"Successfully registered admin"
// @Router			/auth/register [post]
func (con *AuthController) Register() *ctr.Route {
	return ctr.POST("/register").Do(func() []ctr.H {
		body := new(dto.RegisterRequest)

		return []ctr.H{
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.RegisterAdmin(c.Request().Context(), body)
				if err != nil {
					return err
				}

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
func (con *AuthController) Me() *ctr.Route {
	return ctr.GET("/me").Use(con.jwt_s.RequiredAdmin()).Do(func() []ctr.H {
		admin := new(ent.Admin)

		return []ctr.H{
			// extract the admin from context which inject by jwt_s.RequiredAdmin()
			request.MustAdmin(admin),

			// handler
			func(c echo.Context) error {
				return request.Response(c, response.Data(admin))
			},
		}
	})
}
