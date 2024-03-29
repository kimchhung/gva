package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/common/services"
	"github.com/kimchhung/gva/app/module/role/dto"
	"github.com/kimchhung/gva/app/module/role/service"
	"github.com/kimchhung/gva/internal/rctrl"
	"github.com/kimchhung/gva/utils/request"
)

var _ IRoleController = (*RoleController)(nil)

type IRoleController interface {
	rctrl.FiberRouter
	Create(meta *rctrl.RouteMeta) rctrl.MetaHandler
	List(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Get(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Update(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler
}

type RoleController struct {
	service    *service.RoleService
	jwtService *services.JwtService
}

func (con *RoleController) Routes(r fiber.Router) {
	role := r.Group("role")
	role.Use(con.jwtService.ProtectAdmin())
	rctrl.Register(role, con)
}

func NewRoleController(service *service.RoleService, jwtService *services.JwtService) *RoleController {
	return &RoleController{
		service:    service,
		jwtService: jwtService,
	}
}

// @Tags Role
// @Security Bearer
// @Summary List all Roles
// @Description Get a list of all Roles
// @ID list-all-Roles
// @Accept  json
// @Produce  json
// @Success  200 {object} request.Response{data=[]dto.RoleResponse} "Successfully retrieved Roles"
// @Router /role [get]
func (con *RoleController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Roles").Do(func(c *fiber.Ctx) error {
		list, err := con.service.GetRoles(c.UserContext())
		if err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "Role list retreived successfully!",
			Data:    list,
		})
	})
}

// @Tags Role
// @Security Bearer
// @Summary Get a Role
// @Description Get a Role by ID
// @ID get-Role-by-id
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path int true "Role ID"
// @Success   200 {object} request.Response{data=dto.RoleResponse}
// @Router /role/{id} [get]
func (con *RoleController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Role").Do(func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		data, err := con.service.GetRoleByID(c.UserContext(), id)
		if err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "The role retrieved successfully!",
			Data:    data,
		})
	})
}

// @Tags Role
// @Security Bearer
// @Summary Create a Role
// @Description Create a new Role with the provided details
// @ID create-Role
// @Accept  json
// @Produce  json
// @Param Role body dto.RoleRequest true "Role data"
// @Success  200 {object} request.Response{data=dto.RoleResponse} "Successfully created Role"
// @Router /role [post]
func (con *RoleController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Role").DoWithScope(func() []fiber.Handler {
		req := new(dto.RoleRequest)

		return []fiber.Handler{
			request.Validate(request.BodyParser(&req)),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreateRole(c.UserContext(), *req)
				if err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The role was created successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Role
// @Security Bearer
// @Summary Update a Role
// @Description Update a Role by ID
// @ID update-Role-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Role ID"
// @Param Role body dto.RoleRequest true "Role data"
// @Success  200 {object} request.Response{data=dto.RoleResponse} "Successfully updated Role"
// @Router /role/{id} [patch]
func (con *RoleController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Role").DoWithScope(func() []fiber.Handler {
		req := new(dto.RoleRequest)

		return []fiber.Handler{
			request.Validate(request.BodyParser(&req)),
			func(c *fiber.Ctx) error {
				id, err := strconv.Atoi(c.Params("id"))
				if err != nil {
					return err
				}

				data, err := con.service.UpdateRole(c.UserContext(), id, *req)
				if err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The role was updated successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Role
// @Security Bearer
// @Summary Delete a Role
// @Description Delete a Role by ID
// @ID delete-Role-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Role ID"
// @Success  200 {object} request.Response{} "Successfully deleted Role"
// @Router /role/{id} [delete]
func (con *RoleController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Role").Do(func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		if err = con.service.DeleteRole(c.UserContext(), id); err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "The Role was deleted successfully!",
		})
	})
}
