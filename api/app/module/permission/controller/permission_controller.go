package controller

import (
	"strconv"

	"github.com/kimchhung/gva/app/module/permission/dto"
	"github.com/kimchhung/gva/app/module/permission/service"

	"github.com/kimchhung/gva/internal/control_route"
	"github.com/kimchhung/gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = (*PermissionController)(nil)

type PermissionController struct {
	service *service.PermissionService
}

func (con *PermissionController) Routes(r fiber.Router) {
	r.Route(
		"/permissions", func(router fiber.Router) {
			router.Get("/", con.List).Name("get many permissions")
			router.Get("/:id", con.Get).Name("get one permission")
			router.Post("/", con.Create).Name("create one permission")
			router.Patch("/:id", con.Update).Name("update one permission")
			router.Delete("/:id", con.Delete).Name("delete one permission")
		},
	)
}

func NewPermissionController(service *service.PermissionService) *PermissionController {
	return &PermissionController{
		service: service,
	}
}

func (con *PermissionController) List(c *fiber.Ctx) error {
	list, err := con.service.GetPermissions(c.UserContext())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "Permission list retreived successfully!",
		Data:    list,
	})
}

func (con *PermissionController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	data, err := con.service.GetPermissionByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The permission retrieved successfully!",
		Data:    data,
	})
}

func (con *PermissionController) Create(c *fiber.Ctx) error {
	req := new(dto.PermissionRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.CreatePermission(c.UserContext(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The permission was created successfully!",
		Data:    data,
	})
}

func (con *PermissionController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(dto.PermissionRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.UpdatePermission(c.UserContext(), id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The permission was updated successfully!",
		Data:    data,
	})
}

func (con *PermissionController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.service.DeletePermission(c.UserContext(), id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The permission was deleted successfully!",
	})
}
