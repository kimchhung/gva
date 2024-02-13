package module_template

var Controller = `package controller

import (
	"strconv"

	"gva/app/module/{{.EntitySnake}}/dto"
	"gva/app/module/{{.EntitySnake}}/service"

	"gva/internal/control_route"
	"gva/utils/response"

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
