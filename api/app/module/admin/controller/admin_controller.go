package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/common/permissions"
	"github.com/kimchhung/gva/app/common/services"
	"github.com/kimchhung/gva/app/module/admin/dto"
	"github.com/kimchhung/gva/app/module/admin/service"
	"github.com/kimchhung/gva/internal/rctrl"
	"github.com/kimchhung/gva/utils/request"
)

var _ IAdminController = (*AdminController)(nil)

type IAdminController interface {
	rctrl.FiberRouter
	Create(meta *rctrl.RouteMeta) rctrl.MetaHandler
	List(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Get(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Update(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler
}

type AdminController struct {
	service    *service.AdminService
	jwtService *services.JwtService
}

func (con *AdminController) Routes(r fiber.Router) {
	group := r.Group("admin")
	group.Use(
		con.jwtService.ProtectAdmin(),
	)
	rctrl.Register(group, con)
}

func NewAdminController(service *service.AdminService, jwtService *services.JwtService) *AdminController {
	return &AdminController{
		service:    service,
		jwtService: jwtService,
	}
}

// @Tags Admin
// @Summary List all Admins
// @Description Get a list of all Admins
// @ID list-all-Admins
// @Accept  json
// @Produce  json
// @Success  200 {object} request.Response{data=[]dto.AdminResponse} "Successfully retrieved Admins"
// @Router /admin [get]
func (con *AdminController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Admins").Do(
		permissions.RequireAny(
			permissions.AdminView,
			permissions.AdminSuper,
		),
		func(c *fiber.Ctx) error {
			list, err := con.service.GetAdmins(c.UserContext())
			if err != nil {
				return err
			}

			return request.Resp(c, request.Response{
				Message: "Admin list retreived successfully!",
				Data:    list,
			})
		},
	)
}

// @Tags Admin
// @Summary Get a Admin
// @Description Get a Admin by ID
// @ID get-Admin-by-id
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Admin ID"
// @Success   200 {object} request.Response{data=dto.AdminResponse}
// @Router /admin/{id} [get]
func (con *AdminController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Admin").DoWithScope(func() []fiber.Handler {
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

				return request.Resp(c, request.Response{
					Message: "The admin retrieved successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Admin
// @Summary Create a Admin
// @Description Create a new Admin with the provided details
// @ID create-Admin
// @Accept  json
// @Produce  json
// @Param Admin body dto.AdminRequest true "Admin data"
// @Success  200 {object} request.Response{data=dto.AdminResponse} "Successfully created Admin"
// @Router /admin [post]
func (con *AdminController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Admin").DoWithScope(func() []fiber.Handler {
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

				return request.Resp(c, request.Response{
					Message: "The admin was created successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Admin
// @Summary Update a Admin
// @Description Update a Admin by ID
// @ID update-Admin-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Param Admin body dto.AdminRequest true "Admin data"
// @Success  200 {object} request.Response{data=dto.AdminResponse} "Successfully updated Admin"
// @Router /admin/{id} [patch]
func (con *AdminController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Admin").DoWithScope(func() []fiber.Handler {
		req := new(dto.AdminRequest)
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
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

				return request.Resp(c, request.Response{
					Message: "The admin was updated successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Admin
// @Summary Delete a Admin
// @Description Delete a Admin by ID
// @ID delete-Admin-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Success  200 {object} request.Response{} "Successfully deleted Admin"
// @Router /admin/{id} [delete]
func (con *AdminController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Admin").DoWithScope(func() []fiber.Handler {
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

				return request.Resp(c, request.Response{
					Message: "The Admin was deleted successfully!",
				})
			},
		}
	})
}
