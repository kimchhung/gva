package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/app/module/role/dto"
	"github.com/kimchhung/gva/extra/app/module/role/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/utils/request"
)

// don't remove for runtime type checking
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
	service *service.RoleService
}

func (con *RoleController) Routes(r fiber.Router) {
	role := r.Group("role")
	rctrl.Register(role, con)
}

func NewRoleController(service *service.RoleService) *RoleController {
	return &RoleController{
		service: service,
	}
}

// @Tags Role
// @Summary List all Roles
// @Description Get a list of all Roles
// @ID list-all-Roles
// @Accept  json
// @Produce  json
// @Success 200 {object} request.Response{data=map[string]dto.RoleResponse{list=[]dto.RoleResponse}} "Successfully retrieved Roles"
// @Router /role [get]
// @Security Bearer
func (con *RoleController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Roles").Do(func(c *fiber.Ctx) error {
		list, err := con.service.GetRoles(c.UserContext())
		if err != nil {
			return err
		}

		return request.Resp(c,
			request.Data(list),
			request.Message("Role list retreived successfully!"),
		)
	})
}

// @Tags Role
// @Security Bearer
// @Summary Get a Role
// @Description Get a Role by ID
// @ID get-Role-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Role ID"
// @Success   200 {object} request.Response{data=dto.RoleResponse}
// @Router /role/{id} [get]
func (con *RoleController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Role").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.GetRoleByID(c.UserContext(), param.ID)
				if err != nil {
					return err
				}

				return request.Resp(c,
					request.Data(data),
					request.Message("The role retrieved successfully!"),
				)
			},
		}
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
		body := new(dto.RoleRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreateRole(c.UserContext(),body)
				if err != nil {
					return err
				}

				return request.Resp(c,
					request.Data(data),
					request.Message("The role retrieved successfully!"),
				)
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
		body := new(dto.RoleRequest)
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.UpdateRole(c.UserContext(), param.ID,body)
				if err != nil {
					return err
				}

				return request.Resp(c,
					request.Data(data),
					request.Message("The role retrieved successfully!"),
				)
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
func (con  *RoleController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Role").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.DeleteRole(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Resp(c,
					request.Message("The role retrieved successfully!"),
				)
			},
		}
	})
}
