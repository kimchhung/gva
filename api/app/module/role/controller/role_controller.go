package controller

import (
	"strconv"

	"gva/app/module/role/dto"
	"gva/app/module/role/service"

	"gva/internal/control_route"
	"gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = (*RoleController)(nil)

type RoleController struct {
	service *service.RoleService
}

func (con *RoleController) Routes(r fiber.Router) {
	r.Route(
		"/roles", func(router fiber.Router) {
			router.Get("/", con.List).Name("get many roles")
			router.Get("/:id", con.Get).Name("get one role")
			router.Post("/", con.Create).Name("create one role")
			router.Patch("/:id", con.Update).Name("update one role")
			router.Delete("/:id", con.Delete).Name("delete one role")
		},
	)
}

func NewRoleController(service *service.RoleService) *RoleController {
	return &RoleController{
		service: service,
	}
}

func (con *RoleController) List(c *fiber.Ctx) error {
	list, err := con.service.GetRoles(c.UserContext())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "Role list retreived successfully!",
		Data:    list,
	})
}

func (con *RoleController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	data, err := con.service.GetRoleByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The role retrieved successfully!",
		Data:    data,
	})
}

func (con *RoleController) Create(c *fiber.Ctx) error {
	req := new(dto.RoleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.CreateRole(c.UserContext(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The role was created successfully!",
		Data:    data,
	})
}

func (con *RoleController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(dto.RoleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.UpdateRole(c.UserContext(), id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The role was updated successfully!",
		Data:    data,
	})
}

func (con *RoleController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.service.DeleteRole(c.UserContext(), id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The role was deleted successfully!",
	})
}

