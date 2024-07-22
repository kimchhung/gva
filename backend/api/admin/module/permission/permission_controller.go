package permission

import (
	"github.com/gva/api/admin/module/permission/dto"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/echoc"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/internal/rql"
	"github.com/labstack/echo/v4"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*PermissionController)(nil)

type PermissionController struct {
	service *PermissionService
}

func (con *PermissionController) Init(r *echo.Group) *echo.Group {
	return r.Group("/permission")
}

func NewPermissionController(service *PermissionService) *PermissionController {
	return &PermissionController{
		service: service,
	}
}

// @Tags Permission
// @Summary List all Permissions
// @Description Get a list of all Permissions
// @ID list-all-Permissions
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.PermissionResponse{list=[]dto.PermissionResponse}}"
// @Router /permission [get]
// @Security Bearer
func (con *PermissionController) List(meta *echoc.RouteMeta) echoc.MetaHandler {
	parser := request.MustRqlParser(rql.Config{
		Model:        ent.Permission{},
		DefaultLimit: 25,
		DefaultSort:  []string{"-id"},
		FieldSep:     ".",
	})

	return meta.Get("/").DoWithScope(func() []echo.HandlerFunc {
		params := new(dto.PermissionPagedRequest)
		return []echo.HandlerFunc{
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.service.GetPermissions(c.Request().Context(), params)
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

// @Tags Permission
// @Security Bearer
// @Summary Get a Permission
// @Description Get a Permission by ID
// @ID get-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Success   200 {object} response.Response{data=dto.PermissionResponse}
// @Router /permission/{id} [get]
func (con *PermissionController) Get(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/:id").Name("get one Permission").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				data, err := con.service.GetPermissionByID(c.Request().Context(), param.ID)
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

// @Tags Permission
// @Security Bearer
// @Summary Create a Permission
// @Description Create a new Permission with the provided details
// @ID create-Permission
// @Accept  json
// @Produce  json
// @Param Permission body dto.PermissionRequest true "Permission data"
// @Success  200 {object} response.Response{data=dto.PermissionResponse}
// @Router /permission [post]
func (con *PermissionController) Create(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Post("/").Name("create one Permission").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.PermissionRequest)

		return []echo.HandlerFunc{
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

// @Tags Permission
// @Security Bearer
// @Summary Update a Permission
// @Description Update a Permission by ID
// @ID update-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Param Permission body dto.PermissionRequest true "Permission data"
// @Success  200 {object} response.Response{data=dto.PermissionResponse}
// @Router /permission/{id} [patch]
func (con *PermissionController) Update(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Patch("/:id").Name("update one Permission").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.PermissionRequest)
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePermission(c.Request().Context(), param.ID, body)
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

// @Tags Permission
// @Security Bearer
// @Summary Delete a Permission
// @Description Delete a Permission by ID
// @ID delete-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Success  200 {object} response.Response{} "The permission deleted successfully!"
// @Router /permission/{id} [delete]
func (con *PermissionController) Delete(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Delete("/:id").Name("delete one Permission").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				if err := con.service.DeletePermission(c.Request().Context(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The permission deleted successfully!"),
				)
			},
		}
	})
}
