package module_template

import (
	"strings"
)

var Controller = strings.ReplaceAll(`package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/app/module/{{.EntitySnake}}/dto"
	"github.com/kimchhung/gva/extra/app/module/{{.EntitySnake}}/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/utils/request"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*{{.EntityPascal}}Controller)(nil)


type {{.EntityPascal}}Controller struct {
	service *service.{{.EntityPascal}}Service
}

func (con *{{.EntityPascal}}Controller) Routes(r fiber.Router) {
	{{.EntityAllLower}} := r.Group("{{.EntityKebab}}")
	rctrl.Register({{.EntityAllLower}}, con)
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
// @Success 200 {object} request.Response{data=map[string]dto.{{.EntityPascal}}Response{list=[]dto.{{.EntityPascal}}Response}} "Successfully retrieved {{.EntityPascal}}s"
// @Router /{{.EntityKebab}} [get]
// @Security Bearer
func (con *{{.EntityPascal}}Controller) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many {{.EntityPascal}}s").Do(func(c *fiber.Ctx) error {
		list, err := con.service.Get{{.EntityPascal}}s(c.UserContext())
		if err != nil {
			return err
		}

		return request.Resp(c,
			request.Data(list),
			request.Message("{{.EntityPascal}} list retreived successfully!"),
		)
	})
}

// @Tags {{.EntityPascal}}
// @Security Bearer
// @Summary Get a {{.EntityPascal}}
// @Description Get a {{.EntityPascal}} by ID
// @ID get-{{.EntityPascal}}-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "{{.EntityPascal}} ID"
// @Success   200 {object} request.Response{data=dto.{{.EntityPascal}}Response}
// @Router /{{.EntityKebab}}/{id} [get]
func (con *{{.EntityPascal}}Controller) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one {{.EntityPascal}}").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int %sparams:"id" validate:"gt=0"%s
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.Get{{.EntityPascal}}ByID(c.UserContext(), param.ID)
				if err != nil {
					return err
				}

				return request.Resp(c,
					request.Data(data),
					request.Message("The {{.EntityAllLower}} retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags {{.EntityPascal}}
// @Security Bearer
// @Summary Create a {{.EntityPascal}}
// @Description Create a new {{.EntityPascal}} with the provided details
// @ID create-{{.EntityPascal}}
// @Accept  json
// @Produce  json
// @Param {{.EntityPascal}} body dto.{{.EntityPascal}}Request true "{{.EntityPascal}} data"
// @Success  200 {object} request.Response{data=dto.{{.EntityPascal}}Response} "Successfully created {{.EntityPascal}}"
// @Router /{{.EntityKebab}} [post]
func (con *{{.EntityPascal}}Controller) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one {{.EntityPascal}}").DoWithScope(func() []fiber.Handler {
		body := new(dto.{{.EntityPascal}}Request)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				data, err := con.service.Create{{.EntityPascal}}(c.UserContext(),body)
				if err != nil {
					return err
				}

				return request.Resp(c,
					request.Data(data),
					request.Message("The {{.EntityAllLower}} retrieved successfully!"),
				)
			},
		}
	})
}


// @Tags {{.EntityPascal}}
// @Security Bearer
// @Summary Update a {{.EntityPascal}}
// @Description Update a {{.EntityPascal}} by ID
// @ID update-{{.EntityPascal}}-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "{{.EntityPascal}} ID"
// @Param {{.EntityPascal}} body dto.{{.EntityPascal}}Request true "{{.EntityPascal}} data"
// @Success  200 {object} request.Response{data=dto.{{.EntityPascal}}Response} "Successfully updated {{.EntityPascal}}"
// @Router /{{.EntityKebab}}/{id} [patch]
func (con *{{.EntityPascal}}Controller) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one {{.EntityPascal}}").DoWithScope(func() []fiber.Handler {
		body := new(dto.{{.EntityPascal}}Request)
		param := &struct {
			ID int %sparams:"id" validate:"gt=0"%s
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.Update{{.EntityPascal}}(c.UserContext(), param.ID,body)
				if err != nil {
					return err
				}

				return request.Resp(c,
					request.Data(data),
					request.Message("The {{.EntityAllLower}} retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags {{.EntityPascal}}
// @Security Bearer
// @Summary Delete a {{.EntityPascal}}
// @Description Delete a {{.EntityPascal}} by ID
// @ID delete-{{.EntityPascal}}-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "{{.EntityPascal}} ID"
// @Success  200 {object} request.Response{} "Successfully deleted {{.EntityPascal}}"
// @Router /{{.EntityKebab}}/{id} [delete]
func (con  *{{.EntityPascal}}Controller) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one {{.EntityPascal}}").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int %sparams:"id" validate:"gt=0"%s
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.Delete{{.EntityPascal}}(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Resp(c,
					request.Message("The {{.EntityAllLower}} retrieved successfully!"),
				)
			},
		}
	})
}
`,
	"%s", special,
)
