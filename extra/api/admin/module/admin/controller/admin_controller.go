package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/api/admin/module/admin/dto"
	"github.com/kimchhung/gva/extra/api/admin/module/admin/service"
	"github.com/kimchhung/gva/extra/app/common/contexts"
	"github.com/kimchhung/gva/extra/app/common/permissions"
	"github.com/kimchhung/gva/extra/app/common/services"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/internal/rql"
)

var _ interface{ rctrl.Controller } = (*AdminController)(nil)

type AdminController struct {
	service    *service.AdminService
	jwtService *services.JwtService
}

func (con *AdminController) Init(r fiber.Router) fiber.Router {
	return r.Group("admin").Use(
		con.jwtService.ProtectAdmin(),
	)
}

func NewAdminController(service *service.AdminService, jwtService *services.JwtService) *AdminController {
	return &AdminController{
		service:    service,
		jwtService: jwtService,
	}
}

// @Tags Admin
// @Security Bearer
// @Summary List all Admins
// @Description Get a list of all Admins
// @ID list-all-Admins
// @Accept  json
// @Produce  json
// @Success  200 {object} response.Response{data=[]ent.Admin,meta=pagi.Meta} "Successfully retrieved Admins"
// @Router /admin [get]
func (con *AdminController) Paginate(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	// init parser once and reused
	parser := request.MustRqlParser("admins",
		struct{ ent.Admin }{},
	)

	return meta.Get("/paginate").DoWithScope(func() []fiber.Handler {
		params := new(rql.Params)

		return []fiber.Handler{
			permissions.RequireAny(
				permissions.AdminView,
				permissions.AdminSuper,
			),
			request.Parse(
				request.RqlParser(params, parser),
			),
			func(c *fiber.Ctx) error {
				list, meta, err := con.service.Paginate(c.UserContext(), params)
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

// @Tags Admin
// @Security Bearer
// @Summary Get Admin Routes
// @Description Get a list of routes for an Admin by ID
// @ID get-Admin-routes
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{} "Successfully retrieved Admin routes"
// @Router /admin/route [get]
func (con *AdminController) AdminRoutes(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/route").DoWithScope(func() []fiber.Handler {
		adminCtx := new(contexts.AdminContext)

		return []fiber.Handler{
			request.AdminContext(adminCtx),
			func(c *fiber.Ctx) error {
				list, err := con.service.GetAdminNestedRouteById(c.UserContext(), adminCtx.Admin.ID)
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

// @Tags Admin
// @Security Bearer
// @Summary Get Admin permissions
// @Description Get a list of permissions for an Admin by ID
// @ID get-Admin-permissions
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{} "Successfully retrieved Admin permissions"
// @Router /admin/route [get]
func (con *AdminController) AdminPermissions(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/permission").DoWithScope(func() []fiber.Handler {
		var admin *ent.Admin

		return []fiber.Handler{
			request.Admin(admin),
			func(c *fiber.Ctx) error {
				permissions, err := con.service.GetAdminPermissionById(c.UserContext(), admin.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(permissions),
				)
			},
		}
	})
}

// @Tags Admin
// @Security Bearer
// @Summary Get a Admin
// @Description Get a Admin by ID
// @ID get-Admin-by-id
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path int true "Admin ID"
// @Success   200 {object} response.Response{data=dto.AdminResponse}
// @Router /admin/{id} [get]
func (con *AdminController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
		}{}

		return []fiber.Handler{
			permissions.RequireAny(
				permissions.AdminView,
				permissions.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.GetAdminByID(c.UserContext(), param.ID)
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

// @Tags Admin
// @Security Bearer
// @Summary Create a Admin
// @Description Create a new Admin with the provided details
// @ID create-Admin
// @Accept  json
// @Produce  json
// @Param Admin body dto.AdminRequest true "Admin data"
// @Success  200 {object} response.Response{data=dto.AdminResponse} "Successfully created Admin"
// @Router /admin [post]
func (con *AdminController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").DoWithScope(func() []fiber.Handler {
		req := new(dto.AdminRequest)

		return []fiber.Handler{
			permissions.RequireAny(
				permissions.AdminModify,
				permissions.AdminSuper,
			),
			request.Validate(
				request.BodyParser(req),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.CreateAdmin(c.UserContext(), *req)
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

// @Tags Admin
// @Security Bearer
// @Summary Update a Admin
// @Description Update a Admin by ID
// @ID update-Admin-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Param Admin body dto.AdminRequest true "Admin data"
// @Success  200 {object} response.Response{data=dto.AdminResponse} "Successfully updated Admin"
// @Router /admin/{id} [patch]
func (con *AdminController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").DoWithScope(func() []fiber.Handler {
		req := new(dto.AdminRequest)
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			permissions.RequireAny(
				permissions.AdminModify,
				permissions.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(req),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.UpdateAdmin(c.UserContext(), param.ID, *req)
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

// @Tags Admin
// @Security Bearer
// @Summary Delete a Admin
// @Description Delete a Admin by ID
// @ID delete-Admin-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Success  200 {object} response.Response{} "Successfully deleted Admin"
// @Router /admin/{id} [delete]
func (con *AdminController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
		}{}

		return []fiber.Handler{
			permissions.RequireAny(
				permissions.AdminDelete,
				permissions.AdminSuper,
			),
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.DeleteAdmin(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The Admin was deleted successfully!"),
				)
			},
		}
	})
}
