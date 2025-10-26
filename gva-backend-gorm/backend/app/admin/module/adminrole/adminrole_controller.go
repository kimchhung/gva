package adminrole

import (
	"backend/app/admin/middleware"
	"backend/app/admin/module/adminrole/dto"
	"backend/app/share/permission"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*AdminRoleController)(nil)

type AdminRoleController struct {
	middleware *middleware.Middleware
	service    *AdminRoleService
}

func NewAdminRoleController(
	middleware *middleware.Middleware,
	service *AdminRoleService,
) *AdminRoleController {
	return &AdminRoleController{
		middleware: middleware,
		service:    service,
	}
}

func (con *AdminRoleController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/adminrole",
			con.middleware.JwtGuard(),
			con.middleware.IpGuard(),
		),
	)
}

// @Tags			AdminRole
// @Summary			Create AdminRole
// @Description		Create AdminRole
// @ID				create-adminrole
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.CreateAdminRoleRequest	true	"AdminRole data"
// @Success			200		{object}	response.Response{data=dto.AdminRoleResponse}	"Successfully created AdminRole"
// @Router			/adminrole/ [post]
func (con *AdminRoleController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.CreateAdminRoleRequest)

		return []ctr.H{
			permission.RequireAny(permission.AdminRoleAdd),
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.CreateAdminRole(c.Request().Context(), body)
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

// @Tags			AdminRole
// @Summary			Get AdminRole
// @Description		Get AdminRole
// @ID				get-adminrole
// @Accept			json
// @Param			id	path	int	true	"AdminRole ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.AdminRoleResponse}	"Successfully get AdminRole"
// @Router			/adminrole/{id} [get]
func (con *AdminRoleController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRoleRequest)

		return []ctr.H{
			permission.RequireAny(permission.AdminRoleView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetAdminRole(c.Request().Context(), params.ID)
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

// @Tags			AdminRole
// @Summary			Get AdminRoles
// @Description		Get AdminRoles
// @ID				get-adminroles
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
// @Success			200		{object}	response.Response{data=[]dto.AdminRoleResponse}	"Successfully get AdminRoles"
// @Router			/adminrole/ [get]
func (con *AdminRoleController) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAny(permission.AdminRoleView),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.GetAdminRoles(c.Request().Context(), query)
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

// @Tags			AdminRole
// @Summary			Update AdminRole
// @Description		Update AdminRole
// @ID				update-adminrole
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"AdminRole ID"
// @Param			body	body		dto.UpdateAdminRoleRequest	true	"AdminRole data"
// @Success			200		{object}	response.Response{data=dto.AdminRoleResponse}	"Successfully updated AdminRole"
// @Router			/adminrole/{id} [put]
func (con *AdminRoleController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRoleRequest)
		body := new(dto.UpdateAdminRoleRequest)

		return []ctr.H{
			permission.RequireAny(permission.AdminRoleEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateAdminRole(c.Request().Context(), params.ID, body)
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

// @Tags			AdminRole
// @Summary			Update AdminRole partial
// @Description		Update AdminRole partial
// @ID				update-adminrole-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "AdminRole ID"
// @Param			body	body	dto.UpdateAdminRoleRequest	true	"AdminRole data"
// @Success			200		{object}	response.Response{data=dto.AdminRoleResponse}	"Successfully updated AdminRole"
// @Router			/adminrole/{id} [patch]
func (con *AdminRoleController) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRoleRequest)
		body := new(dto.UpdatePatchAdminRoleRequest)

		return []ctr.H{
			permission.RequireAny(permission.AdminRoleEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePatchAdminRole(c.Request().Context(), params.ID, body)
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

// @Tags			AdminRole
// @Summary			Delete AdminRole
// @Description		Delete AdminRole
// @ID				delete-adminrole
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"AdminRole ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted AdminRole"
// @Router			/adminrole/{id} [delete]
func (con *AdminRoleController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRoleRequest)

		return []ctr.H{
			permission.RequireAny(permission.AdminRoleDelete),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeleteAdminRole(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted AdminRole"),
				)
			},
		}
	})
}
