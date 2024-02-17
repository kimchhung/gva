package module_template

var Controller = `package controller

import (
	"strconv"

	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/dto"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/service"

	"github.com/kimchhung/gva/internal/control_route"
	"github.com/kimchhung/gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = (*{{.Entity}}Controller)(nil)

type {{.Entity}}Controller struct {
	service *service.{{.Entity}}Service
}

func (con *{{.Entity}}Controller) Routes(r fiber.Router) {
	r.Route(
		"/{{.EntityLower}}s", func(router fiber.Router) {
			router.Get("/", con.List).Name("get many {{.EntityLower}}s")
			router.Get("/:id", con.Get).Name("get one {{.EntityLower}}")
			router.Post("/", con.Create).Name("create one {{.EntityLower}}")
			router.Patch("/:id", con.Update).Name("update one {{.EntityLower}}")
			router.Delete("/:id", con.Delete).Name("delete one {{.EntityLower}}")
		},
	)
}

func New{{.Entity}}Controller(service *service.{{.Entity}}Service) *{{.Entity}}Controller {
	return &{{.Entity}}Controller{
		service: service,
	}
}

// @Summary List all {{.Entity}}s
// @Description Get a list of all {{.Entity}}s
// @Tags {{.Entity}}
// @Accept  json
// @Produce  json
// @Success  200 {object} response.Response{Data=[]any}
// @Router /{{.Entity}}s [get]
func (con *{{.Entity}}Controller) List(c *fiber.Ctx) error {
	list, err := con.service.Get{{.Entity}}s(c.UserContext())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "{{.Entity}} list retreived successfully!",
		Data:    list,
	})
}

// Get godoc
// @Summary Get one {{.Entity}} by ID
// @Description Get details of an {{.Entity}} by ID
// @Tags {{.Entity}}
// @Accept  json
// @Produce  json
// @Param id path int true "{{.Entity}} ID"
// @Success  200 {object} response.Response{}
// @Router /{{.Entity}}s/{id} [get]
func (con *{{.Entity}}Controller) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	data, err := con.service.Get{{.Entity}}ByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The {{.EntitySnake}} retrieved successfully!",
		Data:    data,
	})
}

// Create godoc
// @Summary Create a new {{.Entity}}
// @Description Add a new {{.Entity}} to the system
// @Tags {{.Entity}}
// @Accept  json
// @Produce  json
// @Param {{.Entity}} body dto.{{.Entity}}Request true "{{.Entity}} data"
// @Success  200 {object} response.Response{}
// @Router /{{.Entity}}s [post]
func (con *{{.Entity}}Controller) Create(c *fiber.Ctx) error {
	req := new(dto.{{.Entity}}Request)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.Create{{.Entity}}(c.UserContext(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The {{.EntitySnake}} was created successfully!",
		Data:    data,
	})
}

// Update godoc
// @Summary Update an existing {{.Entity}}
// @Description Update the details of an {{.Entity}} by ID
// @Tags {{.Entity}}
// @Accept  json
// @Produce  json
// @Param id path int true "{{.Entity}} ID"
// @Param {{.Entity}} body dto.{{.Entity}}Request true "{{.Entity}} data"
// @Success  200 {object} response.Response{}
// @Router /{{.Entity}}s/{id} [patch]
func (con *{{.Entity}}Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(dto.{{.Entity}}Request)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.Update{{.Entity}}(c.UserContext(), id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The {{.EntitySnake}} was updated successfully!",
		Data:    data,
	})
}

// Delete godoc
// @Summary Delete an {{.Entity}}
// @Description Delete an {{.Entity}} by ID
// @Tags {{.Entity}}
// @Accept  json
// @Produce  json
// @Param id path int true "{{.Entity}} ID"
// @Success  200 {object} response.Response{}
// @Router /{{.Entity}}s/{id} [delete]
func (con *{{.Entity}}Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.service.Delete{{.Entity}}(c.UserContext(), id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The {{.EntitySnake}} was deleted successfully!",
	})
}


`
