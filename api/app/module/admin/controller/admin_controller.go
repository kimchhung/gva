package controller

import (
	"strconv"

	"gva/app/module/admin/request"
	"gva/app/module/admin/service"
	"gva/internal/control_route"
	"gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = &AdminController{}

type AdminController struct {
	AdminService *service.AdminService
}

func NewAdminController(AdminService *service.AdminService) *AdminController {
	return &AdminController{
		AdminService: AdminService,
	}
}

func (con *AdminController) Routes(r fiber.Router) {
	r.Route(
		"/admins", func(router fiber.Router) {
			router.Get("/", con.List)
			router.Get("/:id", con.Get)
			router.Post("/", con.Create)
			router.Patch("/:id", con.Update)
			router.Delete("/:id", con.Destroy)
		},
	)
}

func (con *AdminController) List(c *fiber.Ctx) error {
	Admins, err := con.AdminService.GetAdmins(c.UserContext())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "Admin list retreived successfully!",
		Data:    Admins,
	})
}

func (con *AdminController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	Admin, err := con.AdminService.GetAdminByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The Admin retrieved successfully!",
		Data:    Admin,
	})
}

func (con *AdminController) Create(c *fiber.Ctx) error {
	req := new(request.AdminRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	Admin, err := con.AdminService.CreateAdmin(c.UserContext(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The Admin was created successfully!",
		Data:    Admin,
	})
}

func (con *AdminController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(request.AdminRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	Admin, err := con.AdminService.UpdateAdmin(c.UserContext(), id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The Admin was updated successfully!",
		Data:    Admin,
	})
}

func (con *AdminController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.AdminService.DeleteAdmin(c.UserContext(), id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The Admin was deleted successfully!",
	})
}
