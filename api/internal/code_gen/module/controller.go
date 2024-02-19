package module_template

var Controller = `package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/dto"
	"github.com/kimchhung/gva/app/module/{{.EntitySnake}}/service"
	"github.com/kimchhung/gva/internal/rctrl"
	"github.com/kimchhung/gva/utils/response"
)

// don't remove for runtime type checking
var _ I{{.EntityPascal}}Controller = (*{{.EntityPascal}}Controller)(nil)

type I{{.EntityPascal}}Controller interface {
	rctrl.FiberRouter
	Create(meta *rctrl.RouteMeta) rctrl.MetaHandler
	List(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Get(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Update(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler
}


type {{.EntityPascal}}Controller struct {
	service *service.{{.EntityPascal}}Service
}

func (con *{{.EntityPascal}}Controller) Routes(r fiber.Router) {
	r.Route("{{.EntityKebab}}",
		func(router fiber.Router) {
			rctrl.Register(router, con)
		},
	)
}

func New{{.EntityPascal}}Controller(service *service.{{.EntityPascal}}Service) *{{.EntityPascal}}Controller {
	return &{{.EntityPascal}}Controller{
		service: service,
	}
}

// @Tags {{.EntityPascal}}
// @Summary List all {{.EntityPascal}}s
// @Description Get a list of all {{.EntityPascal}}s
// @ID list-all-{{.EntityPascal}}s
// @Accept  json
// @Produce  json
// @Success  200 {object} response.Response{data=[]dto.{{.EntityPascal}}Response} "Successfully retrieved {{.EntityPascal}}s"
// @Router /{{.EntityKebab}} [get]
func (con *{{.EntityPascal}}Controller) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many {{.EntityPascal}}s").Do(func(c *fiber.Ctx) error {
		list, err := con.service.Get{{.EntityPascal}}s(c.UserContext())
		if err != nil {
			return err
		}

		return response.Resp(c, response.Response{
			Message: "{{.EntityPascal}} list retreived successfully!",
			Data:    list,
		})
	})
}

// @Tags {{.EntityPascal}}
// @Summary Get a {{.EntityPascal}}
// @Description Get a {{.EntityPascal}} by ID
// @ID get-{{.EntityPascal}}-by-id
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "{{.EntityPascal}} ID"
// @Success   200 {object} response.Response{data=dto.{{.EntityPascal}}Response}
// @Router /{{.EntityKebab}}/{id} [get]
func (con *{{.EntityPascal}}Controller) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one {{.EntityPascal}}").Do(func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		data, err := con.service.Get{{.EntityPascal}}ByID(c.UserContext(), id)
		if err != nil {
			return err
		}

		return response.Resp(c, response.Response{
			Message: "The {{.EntityCamel}} retrieved successfully!",
			Data:    data,
		})
	})
}

// @Tags {{.EntityPascal}}
// @Summary Create a {{.EntityPascal}}
// @Description Create a new {{.EntityPascal}} with the provided details
// @ID create-{{.EntityPascal}}
// @Accept  json
// @Produce  json
// @Param {{.EntityPascal}} body dto.{{.EntityPascal}}Request true "{{.EntityPascal}} data"
// @Success  200 {object} response.Response{data=dto.{{.EntityPascal}}Response} "Successfully created {{.EntityPascal}}"
// @Router /{{.EntityKebab}} [post]
func (con *{{.EntityPascal}}Controller) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one {{.EntityPascal}}").DoWithScope(func() []fiber.Handler {
		req := new(dto.{{.EntityPascal}}Request)

		return []fiber.Handler{
			response.MustParseAndValidate(&req),

			func(c *fiber.Ctx) error {
				data, err := con.service.Create{{.EntityPascal}}(c.UserContext(), *req)
				if err != nil {
					return err
				}

				return response.Resp(c, response.Response{
					Message: "The {{.EntityCamel}} was created successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags {{.EntityPascal}}
// @Summary Update a {{.EntityPascal}}
// @Description Update a {{.EntityPascal}} by ID
// @ID update-{{.EntityPascal}}-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "{{.EntityPascal}} ID"
// @Param {{.EntityPascal}} body dto.{{.EntityPascal}}Request true "{{.EntityPascal}} data"
// @Success  200 {object} response.Response{data=dto.{{.EntityPascal}}Response} "Successfully updated {{.EntityPascal}}"
// @Router /{{.EntityKebab}}/{id} [patch]
func (con *{{.EntityPascal}}Controller) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one {{.EntityPascal}}").DoWithScope(func() []fiber.Handler {
		req := new(dto.{{.EntityPascal}}Request)

		return []fiber.Handler{
			response.MustParseAndValidate(&req),
			func(c *fiber.Ctx) error {
				id, err := strconv.Atoi(c.Params("id"))
				if err != nil {
					return err
				}

				data, err := con.service.Update{{.EntityPascal}}(c.UserContext(), id, *req)
				if err != nil {
					return err
				}

				return response.Resp(c, response.Response{
					Message: "The {{.EntityCamel}} was updated successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags {{.EntityPascal}}
// @Summary Delete a {{.EntityPascal}}
// @Description Delete a {{.EntityPascal}} by ID
// @ID delete-{{.EntityPascal}}-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "{{.EntityPascal}} ID"
// @Success  200 {object} response.Response{} "Successfully deleted {{.EntityPascal}}"
// @Router /{{.EntityKebab}}/{id} [delete]
func (con *{{.EntityPascal}}Controller) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one {{.EntityPascal}}").Do(func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		if err = con.service.Delete{{.EntityPascal}}(c.UserContext(), id); err != nil {
			return err
		}

		return response.Resp(c, response.Response{
			Message: "The {{.EntityPascal}} was deleted successfully!",
		})
	})
}
`
