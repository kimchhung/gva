package code_gen

var contoller_template = `package controller

import (
	"strconv"

	"gva/app/module/{{.EntitySnake}}/request"
	"gva/app/module/{{.EntitySnake}}/service"

	"gva/internal/control_route"
	"gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = &{{.Entity}}Controller{}

type {{.Entity}}Controller struct {
	service *service.{{.Entity}}Service
}

func (con *{{.Entity}}Controller) Routes(r fiber.Router) {
	r.Route(
		"/{{.EntityLower}}s", func(router fiber.Router) {
			router.Get("/", con.List)
			router.Get("/:id", con.Get)
			router.Post("/", con.Create)
			router.Patch("/:id", con.Update)
			router.Delete("/:id", con.Destroy)
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
	req := new(request.{{.Entity}}Request)
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

	req := new(request.{{.Entity}}Request)
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

func (con *{{.Entity}}Controller) Destroy(c *fiber.Ctx) error {
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
