package admin

import (
	"backend/api/admin/module/admin/dto"
	"backend/app/common/permission"
	"backend/app/common/service"
	"backend/internal/ctr"
	"backend/internal/request"
	"backend/internal/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*AdminController)(nil)

type AdminController struct {
	service *AdminService
	jwt_s   *service.JwtService
	ip_s    *service.IPService
}

func NewAdminController(service *AdminService, jwt_s *service.JwtService, ip_s *service.IPService) *AdminController {
	return &AdminController{
		service: service,
		jwt_s:   jwt_s,
		ip_s:    ip_s,
	}
}

func (con *AdminController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/admin", con.jwt_s.RequiredAdmin(), con.ip_s.RequiredWhiteListIP()),
	)
}

// @Tags			Admin
// @Summary			Create Admin
// @Description		Create Admin
// @ID				create-admin
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.CreateAdminRequest	true	"Admin data"
// @Success			200		{object}	response.Response{data=dto.AdminResponse}	"Successfully created Admin"
// @Router			/admin/ [post]
func (con *AdminController) Create() *ctr.Route {
	return ctr.POST("/").Do(
		func() []ctr.H {
			body := new(dto.CreateAdminRequest)

			return []ctr.H{
				permission.RequireAnyWithSuper(permission.AdminAdd),
				request.Validate(
					request.BodyParser(body),
				),
				func(c echo.Context) error {
					data, err := con.service.CreateAdmin(c.Request().Context(), body)
					if err != nil {
						return err
					}

					return request.Response(c,
						response.Data(data),
					)
				},
			}
		})
}

// @Tags			Admin
// @Summary			Get Admin
// @Description		Get Admin
// @ID				get-admin
// @Accept			json
// @Param			id	path	int	true	"Admin ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.AdminResponse}	"Successfully get Admin"
// @Router			/admin/{id} [get]
func (con *AdminController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetAdmin(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Admin
// @Summary			Get Admins
// @Description		Get Admins
// @ID				get-admins
// @Accept			json
// @Param 			page	query	int	false	"page"
// @Param 			limit	query	int	false	"limit"
// @Param 			search	query	string	false	"search"
// @Param 			filters	query	string	false	"filters"
// @Param 			sorts	query	string	false	"sorts"
// @Param 			orders	query	string	false	"orders"
// @Param 			selects	query	string	false	"selects: list, totalCount"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=[]dto.AdminResponse}	"Successfully get Admins"
// @Router			/admin/ [get]
func (con *AdminController) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminView),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.GetAdmins(c.Request().Context(), query)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(list),
					response.Meta(meta),
				)
			},
		}
	})
}

// @Tags			Admin
// @Summary			Update Admin
// @Description		Update Admin
// @ID				update-admin
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Admin ID"
// @Param			body	body		dto.UpdateAdminRequest	true	"Admin data"
// @Success			200		{object}	response.Response{data=dto.AdminResponse}	"Successfully updated Admin"
// @Router			/admin/{id} [put]
func (con *AdminController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRequest)
		body := new(dto.UpdateAdminRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateAdmin(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Admin
// @Summary			Update Admin partial
// @Description		Update Admin partial
// @ID				update-admin-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "Admin ID"
// @Param			body	body	dto.UpdateAdminRequest	true	"Admin data"
// @Success			200		{object}	response.Response{data=dto.AdminResponse}	"Successfully updated Admin"
// @Router			/admin/{id} [patch]
func (con *AdminController) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRequest)
		body := new(dto.UpdatePatchAdminRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePatchAdmin(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Admin
// @Summary			Set Admin TOTP
// @Description		Set Admin TOTP
// @ID				set-admin-totp
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "Admin ID"
// @Param			body	body	dto.SetTOTPAdminRequest	true	"Admin data"
// @Success			200		{object}	response.Response{data=dto.SetTOTPAdminResponse}	"Successfully set Admin totp"
// @Router			/admin/{id}/totp [patch]
func (con *AdminController) SetAdminTOTP() *ctr.Route {
	return ctr.PATCH("/:id/totp").Do(func() []ctr.H {
		params := new(dto.GetAdminRequest)
		body := new(dto.SetTOTPAdminRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.SetAdminTOTP(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Admin
// @Summary			Delete Admin
// @Description		Delete Admin
// @ID				delete-admin
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Admin ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted Admin"
// @Router			/admin/{id} [delete]
func (con *AdminController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminDelete),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeleteAdmin(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted Admin"),
				)
			},
		}
	})
}
