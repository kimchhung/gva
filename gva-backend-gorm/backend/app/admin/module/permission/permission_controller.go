package permission

import (
	adminmiddleware "backend/app/admin/middleware"
	"backend/app/admin/module/permission/dto"
	"backend/app/share/service"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*PermissionController)(nil)

type PermissionController struct {
	service *PermissionService
	jwt_s   *service.JwtService
	ip_s    *service.IPService
}

func NewPermissionController(service *PermissionService, jwt_s *service.JwtService, ip_s *service.IPService) *PermissionController {
	return &PermissionController{
		service: service,
		jwt_s:   jwt_s,
		ip_s:    ip_s,
	}
}

func (con *PermissionController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/permission", con.jwt_s.RequiredAdmin(), con.ip_s.RequiredWhiteListIP()),
	)
}

// @Tags			Permission
// @Summary			Create Permission
// @Description		Create Permission
// @ID				create-permission
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.CreatePermissionRequest	true	"Permission data"
// @Success			200		{object}	response.Response{data=dto.PermissionResponse}	"Successfully created Permission"
// @Router			/permission/ [post]
func (con *PermissionController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.CreatePermissionRequest)

		return []ctr.H{
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.CreatePermission(c.Request().Context(), body)
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

// @Tags			Permission
// @Summary			Get Permission
// @Description		Get Permission
// @ID				get-permission
// @Accept			json
// @Param			id	path	int	true	"Permission ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.PermissionResponse}	"Successfully get Permission"
// @Router			/permission/{id} [get]
func (con *PermissionController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetPermissionRequest)

		return []ctr.H{
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetPermission(c.Request().Context(), params.ID)
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

// @Tags			Permission
// @Summary			Get Permissions
// @Description		Get Permissions
// @ID				get-permissions
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
// @Success			200		{object}	response.Response{data=[]dto.PermissionResponse}	"Successfully get Permissions"
// @Router			/permission/ [get]
func (con *PermissionController) GetMany() *ctr.Route {
	return ctr.GET("/").Use(adminmiddleware.SkipOperationLog()).Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{

			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.GetPermissions(c.Request().Context(), query)

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

// @Tags			Permission
// @Summary			Update Permission
// @Description		Update Permission
// @ID				update-permission
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Permission ID"
// @Param			body	body		dto.UpdatePermissionRequest	true	"Permission data"
// @Success			200		{object}	response.Response{data=dto.PermissionResponse}	"Successfully updated Permission"
// @Router			/permission/{id} [put]
func (con *PermissionController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.GetPermissionRequest)
		body := new(dto.UpdatePermissionRequest)

		return []ctr.H{
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePermission(c.Request().Context(), params.ID, body)
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

// @Tags			Permission
// @Summary			Update Permission partial
// @Description		Update Permission partial
// @ID				update-permission-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "Permission ID"
// @Param			body	body	dto.UpdatePermissionRequest	true	"Permission data"
// @Success			200		{object}	response.Response{data=dto.PermissionResponse}	"Successfully updated Permission"
// @Router			/permission/{id} [patch]
func (con *PermissionController) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.GetPermissionRequest)
		body := new(dto.UpdatePermissionRequest)

		return []ctr.H{
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePermission(c.Request().Context(), params.ID, body)
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

// @Tags			Permission
// @Summary			Delete Permission
// @Description		Delete Permission
// @ID				delete-permission
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Permission ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted Permission"
// @Router			/permission/{id} [delete]
func (con *PermissionController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.GetPermissionRequest)

		return []ctr.H{
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeletePermission(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted Permission"),
				)
			},
		}
	})
}
