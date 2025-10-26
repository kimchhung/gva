package auth

import (
	admincontext "backend/app/admin/context"
	"backend/app/admin/middleware"
	"backend/app/admin/module/auth/dto"
	"backend/app/share/service"

	"backend/internal/ctr"

	"backend/core/utils/request"
	"backend/core/utils/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*AuthController)(nil)

type AuthController struct {
	middleware *middleware.Middleware
	service    *AuthService
	ip_s       *service.IPService
}

func NewAuthController(
	service *AuthService,
	middleware *middleware.Middleware,
	jwt_s *service.JwtService,
	ip_s *service.IPService,

) *AuthController {
	return &AuthController{
		middleware: middleware,
		service:    service,
		ip_s:       ip_s,
	}
}

func (con *AuthController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/auth",
			middleware.SkipOperationLog(),
		),
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
	return ctr.POST("/login").Use(middleware.SkipOperationLog()).Do(func() []ctr.H {
		body := new(dto.LoginRequest)

		return []ctr.H{
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {

				data, err := con.service.LoginAdmin(
					c.Request().Context(),
					body,
					con.ip_s.GetCurrentIP(c),
				)
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

// @Tags		Auth
// @Security	Bearer
// @Summary		query your self after login
// @Description	get admin data by token
// @ID			create-Auth-me
// @Accept		json
// @Produce		json
// @Success		200	{object}	response.Response{data=model.Admin}	"Successfully registered admin"
// @Router			/auth/me [get]
// @Security		Bearer
func (con *AuthController) Me() *ctr.Route {
	return ctr.GET("/me").
		Use(
			con.middleware.JwtGuard(),
			con.middleware.IpGuard(),
		).Do(func() []ctr.H {
		return []ctr.H{
			// handler
			func(c echo.Context) error {
				admin := admincontext.MustAdminContext(c.Request().Context()).Admin
				return request.Response(c, response.Data(admin))
			},
		}
	})
}
