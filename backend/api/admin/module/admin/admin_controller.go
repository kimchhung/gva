package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gva/api/admin/module/admin/dto"
	appctx "github.com/gva/app/common/context"
	"github.com/gva/app/common/permission"
	"github.com/gva/app/common/service"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/internal/rql"
)

var _ interface{ ctr.CTR } = (*AdminController)(nil)

type AdminController struct {
	service *AdminService
	jwt_s   *service.JwtService
}

func NewAdminController(service *AdminService, jwt_s *service.JwtService) *AdminController {
	return &AdminController{
		service: service,
		jwt_s:   jwt_s,
	}
}

func (con *AdminController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/admin", con.jwt_s.RequiredAdmin()),
	)
}

// @Tags		Admin
// @Security	Bearer
// @Summary		List all Admins
// @Description	Get a list of all Admins
// @ID			list-all-Admins
// @Accept		json
// @Produce		json
// @Success		200	{object}	response.Response{data=[]ent.Admin,meta=pagi.Meta}	"Successfully retrieved Admins"
// @Router		/admin [get]
func (con *AdminController) Paginate() *ctr.Route {
	parser := request.MustRqlParser(rql.Config{
		Model: struct {
			ID xid.ID `json:"id" rql:"filter,sort"`
		}{},
	})

	return ctr.GET("/").Do(func() []ctr.H {
		var (
			params = new(dto.AdminPaginateRequest)
		)
		return []ctr.H{
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
// @Router		/admin/route [get]
func (con *AdminController) AdminRoutes() *ctr.Route {
	return ctr.GET("/routes").Do(func() []ctr.H {
		adminCtx := new(appctx.AdminContext)

		return []ctr.H{
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
// @Router		/admin/permission [get]
func (con *AdminController) AdminPermission() *ctr.Route {
	return ctr.GET("/permission").Do(func() []ctr.H {
		admin := new(ent.Admin)

		return []ctr.H{
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
// @Router		/admin/{id} [get]
func (con *AdminController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		param := new(struct {
			ID xid.ID `param:"id" validate:"required"`
		})

		return []ctr.H{
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
// @Router		/admin [post]
func (con *AdminController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		req := new(dto.AdminRequest)

		return []ctr.H{
			permission.RequireAny(
				permission.AdminAdd,
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
// @Router		/admin/{id} [patch]
func (con *AdminController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		body := new(dto.AdminRequest)
		param := new(struct {
			ID xid.ID `param:"id" validate:"gt=0"`
		})

		return []ctr.H{
			permission.RequireAny(
				permission.AdminAdd,
				permission.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateAdmin(c.Request().Context(), param.ID, body)
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
// @Router		/admin/{id} [delete]
func (con *AdminController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		param := new(struct {
			ID xid.ID `param:"id" validate:"required"`
		})

		return []ctr.H{
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
