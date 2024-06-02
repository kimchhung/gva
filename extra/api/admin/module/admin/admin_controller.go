package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kimchhung/gva/extra/api/admin/module/admin/dto"
	appctx "github.com/kimchhung/gva/extra/app/common/context"
	"github.com/kimchhung/gva/extra/app/common/permission"
	"github.com/kimchhung/gva/extra/internal/echoc"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/internal/rql"
)

var _ interface{ echoc.Controller } = (*AdminController)(nil)

type AdminController struct {
	service *AdminService
}

func (con *AdminController) Init(r *echo.Group) *echo.Group {
	return r.Group("/admins")
}

func NewAdminController(service *AdminService) *AdminController {
	return &AdminController{
		service: service,
	}
}

// @Tags		Admin
// @Security	Bearer
// @Summary		List all Admins
// @Description	Get a list of all Admins
// @ID			list-all-Admins
// @Accept		json
// @Produce		json
// @Success		200	{object}	response.Response{data=[]ent.Admin,meta=pagi.Meta}	"Successfully retrieved Admins"
// @Router		/admins [get]
func (con *AdminController) Paginate(meta *echoc.RouteMeta) echoc.MetaHandler {
	// init parser once and reused
	parser := request.MustRqlParser(rql.Config{
		// Table:        admin.Table,
		Model:        struct{ ent.Admin }{},
		DefaultLimit: 20,
		DefaultSort:  []string{"-id"},
	})

	return meta.Get("/").DoWithScope(func() []echo.HandlerFunc {
		params := new(dto.AdminPaginateRequest)

		return []echo.HandlerFunc{
			permission.RequireAny(
				permission.AdminView,
				permission.AdminSuper,
			),
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.service.Paginate(c.Request().Context(), params)
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

// @Tags		Admin
// @Security	Bearer
// @Summary		Get Admin Routes
// @Description	Get a list of routes for an Admin by ID
// @ID			get-Admin-routes
// @Accept		json
// @Produce		json
// @Success		200	{object}	response.Response{}	"Successfully retrieved Admin routes"
// @Router		/admins/route [get]
func (con *AdminController) AdminRoutes(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/route").DoWithScope(func() []echo.HandlerFunc {
		adminCtx := new(appctx.AdminContext)

		return []echo.HandlerFunc{
			request.MustAdminContext(adminCtx),
			func(c echo.Context) error {
				list, err := con.service.GetAdminNestedRouteById(c.Request().Context(), adminCtx.Admin.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(list),
					response.Message("Admin routes list retreived successfully!"),
				)
			},
		}
	})
}

// @Tags		Admin
// @Security	Bearer
// @Summary		Get Admin permissionissions
// @Description	Get a list of permissionissions for an Admin by ID
// @ID			get-Admin-permissionissions
// @Accept		json
// @Produce		json
// @Success		200	{object}	response.Response{}	"Successfully retrieved Admin permissionissions"
// @Router		/admins/permission [get]
func (con *AdminController) AdminPermission(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/permission").DoWithScope(func() []echo.HandlerFunc {
		var admin *ent.Admin

		return []echo.HandlerFunc{
			request.MustAdmin(admin),
			func(c echo.Context) error {
				permission, err := con.service.GetAdminPermissionById(c.Request().Context(), admin.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(permission),
				)
			},
		}
	})
}

// @Tags		Admin
// @Security	Bearer
// @Summary		Get a Admin
// @Description	Get a Admin by ID
// @ID			get-Admin-by-id
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Param		id	path		int	true	"Admin ID"
// @Success		200	{object}	response.Response{data=dto.AdminResponse}
// @Router		/admins/{id} [get]
func (con *AdminController) Get(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/:id").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
		}{}

		return []echo.HandlerFunc{
			permission.RequireAny(
				permission.AdminView,
				permission.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				data, err := con.service.GetAdminByID(c.Request().Context(), param.ID)
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

// @Tags		Admin
// @Security	Bearer
// @Summary		Create a Admin
// @Description	Create a new Admin with the provided details
// @ID			create-Admin
// @Accept		json
// @Produce		json
// @Param		Admin	body		dto.AdminRequest							true	"Admin data"
// @Success		200		{object}	response.Response{data=dto.AdminResponse}	"Successfully created Admin"
// @Router		/admins [post]
func (con *AdminController) Create(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Post("/").DoWithScope(func() []echo.HandlerFunc {
		req := new(dto.AdminRequest)

		return []echo.HandlerFunc{
			permission.RequireAny(
				permission.AdminModify,
				permission.AdminSuper,
			),
			request.Validate(
				request.BodyParser(req),
			),
			func(c echo.Context) error {
				data, err := con.service.CreateAdmin(c.Request().Context(), req)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Status(http.StatusCreated),
				)
			},
		}
	})
}

// @Tags		Admin
// @Security	Bearer
// @Summary		Update a Admin
// @Description	Update a Admin by ID
// @ID			update-Admin-by-id
// @Accept		json
// @Produce		json
// @Param		id		path		int											true	"Admin ID"
// @Param		Admin	body		dto.AdminRequest							true	"Admin data"
// @Success		200		{object}	response.Response{data=dto.AdminResponse}	"Successfully updated Admin"
// @Router		/admins/{id} [patch]
func (con *AdminController) Update(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Patch("/:id").DoWithScope(func() []echo.HandlerFunc {
		req := new(dto.AdminRequest)
		param := new(struct {
			ID int `params:"id" validate:"gt=0"`
		})

		return []echo.HandlerFunc{
			permission.RequireAny(
				permission.AdminModify,
				permission.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(req),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateAdmin(c.Request().Context(), param.ID, req)
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

// @Tags		Admin
// @Security	Bearer
// @Summary		Delete a Admin
// @Description	Delete a Admin by ID
// @ID			delete-Admin-by-id
// @Accept		json
// @Produce		json
// @Param		id	path		int					true	"Admin ID"
// @Success		200	{object}	response.Response{}	"Successfully deleted Admin"
// @Router		/admins/{id} [delete]
func (con *AdminController) Delete(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Delete("/:id").DoWithScope(func() []echo.HandlerFunc {
		param := new(struct {
			ID int `params:"id" validate:"gte=0"`
		})

		return []echo.HandlerFunc{
			permission.RequireAny(
				permission.AdminDelete,
				permission.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				if err := con.service.DeleteAdmin(c.Request().Context(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The Admin was deleted successfully!"),
				)
			},
		}
	})
}
