package controller

import (
	"strconv"

	"github.com/kimchhung/gva/app/module/admin/dto"
	"github.com/kimchhung/gva/app/module/admin/service"

	"github.com/kimchhung/gva/internal/control_route"
	"github.com/kimchhung/gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = (*AdminController)(nil)

type AdminController struct {
	service    *service.AdminService
	jwtService *service.JWTService
}

func (con *AdminController) Routes(r fiber.Router) {
	r.Route(
		"/admins",
		func(router fiber.Router) {
			// router.Use(con.jwtService.Protected())

			router.Get("/", con.List).Name("get many admins")
			router.Get("/:id", con.Get).Name("get one admin")
			router.Post("/", con.Create).Name("create one admin")
			router.Patch("/:id", con.Update).Name("update one admin")
			router.Delete("/:id", con.Delete).Name("delete one admin")
		},
	)
}

func NewAdminController(service *service.AdminService, jwtService *service.JWTService) *AdminController {
	return &AdminController{
		service:    service,
		jwtService: jwtService,
	}
}

// @Summary List all admins
// @Description Get a list of all admins
// @Tags admin
// @Accept  json
// @Produce  json
// @Success  200 {object} response.Response{Data=[]any}
// @Router /admins [get]
func (con *AdminController) List(c *fiber.Ctx) error {
	list, err := con.service.GetAdmins(c.UserContext())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "Admin list retreived successfully!",
		Data:    list,
	})
}

// Get godoc
// @Summary Get one admin by ID
// @Description Get details of an admin by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Success  200 {object} response.Response{}
// @Router /admins/{id} [get]
func (con *AdminController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	data, err := con.service.GetAdminByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The admin retrieved successfully!",
		Data:    data,
	})
}

// Create godoc
// @Summary Create a new admin
// @Description Add a new admin to the system
// @Tags admin
// @Accept  json
// @Produce  json
// @Param admin body dto.AdminRequest true "Admin data"
// @Success  200 {object} response.Response{}
// @Router /admins [post]
func (con *AdminController) Create(c *fiber.Ctx) error {
	req := new(dto.AdminRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.CreateAdmin(c.UserContext(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The admin was created successfully!",
		Data:    data,
	})
}

// Update godoc
// @Summary Update an existing admin
// @Description Update the details of an admin by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Param admin body dto.AdminRequest true "Admin data"
// @Success  200 {object} response.Response{}
// @Router /admins/{id} [patch]
func (con *AdminController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(dto.AdminRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.UpdateAdmin(c.UserContext(), id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The admin was updated successfully!",
		Data:    data,
	})
}

// Delete godoc
// @Summary Delete an admin
// @Description Delete an admin by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Success  200 {object} response.Response{}
// @Router /admins/{id} [delete]
func (con *AdminController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.service.DeleteAdmin(c.UserContext(), id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The admin was deleted successfully!",
	})
}
