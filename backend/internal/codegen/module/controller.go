package module_template

import (
	"strings"
)

var Controller = strings.ReplaceAll(`package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/backend/api/admin/module/{{.EntitySnake}}/dto"
	"github.com/kimchhung/gva/backend/api/admin/module/{{.EntitySnake}}/service"
	"github.com/kimchhung/gva/backend/internal/rctrl"
	"github.com/kimchhung/gva/backend/internal/request"
	"github.com/kimchhung/gva/backend/internal/response"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*{{.EntityPascal}}Controller)(nil)


type {{.EntityPascal}}Controller struct {
	service *service.{{.EntityPascal}}Service
}

func (con *{{.EntityPascal}}Controller) Init(r *echo.Group) *echo.Group{
	return r.Group("{{.EntityKebab}}")
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
// @Success 200 {object} response.Response{data=map[string]dto.{{.EntityPascal}}Response{list=[]dto.{{.EntityPascal}}Response}} "Successfully retrieved {{.EntityPascal}}s"
// @Router /{{.EntityKebab}} [get]
// @Security Bearer
func (con *{{.EntityPascal}}Controller) List(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/").Name("get many {{.EntityPascal}}s").Do(func(c echo.Context) error {
		list, err := con.service.Get{{.EntityPascal}}s(c.Request().Context())
		if err != nil {
			return err
		}

		return request.Response(c,
			response.Data(list),
			response.Message("{{.EntityPascal}} list retreived successfully!"),
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
// @Success   200 {object} response.Response{data=dto.{{.EntityPascal}}Response}
// @Router /{{.EntityKebab}}/{id} [get]
func (con *{{.EntityPascal}}Controller) Get(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/:id").Name("get one {{.EntityPascal}}").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID int %sparam:"id" validate:"gt=0"%s
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				data, err := con.service.Get{{.EntityPascal}}ByID(c.Request().Context(), param.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The {{.EntityAllLower}} retrieved successfully!"),
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
// @Success  200 {object} response.Response{data=dto.{{.EntityPascal}}Response} "Successfully created {{.EntityPascal}}"
// @Router /{{.EntityKebab}} [post]
func (con *{{.EntityPascal}}Controller) Create(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Post("/").Name("create one {{.EntityPascal}}").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.{{.EntityPascal}}Request)

		return []echo.HandlerFunc{
			request.Validate(
				request.BodyParser(body),
			),

			func(c echo.Context) error {
				data, err := con.service.Create{{.EntityPascal}}(c.Request().Context(),body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The {{.EntityAllLower}} retrieved successfully!"),
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
// @Success  200 {object} response.Response{data=dto.{{.EntityPascal}}Response} "Successfully updated {{.EntityPascal}}"
// @Router /{{.EntityKebab}}/{id} [patch]
func (con *{{.EntityPascal}}Controller) Update(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Patch("/:id").Name("update one {{.EntityPascal}}").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.{{.EntityPascal}}Request)
		param := &struct {
			ID int %sparam:"id" validate:"gt=0"%s
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.Update{{.EntityPascal}}(c.Request().Context(), param.ID,body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The {{.EntityAllLower}} retrieved successfully!"),
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
// @Success  200 {object} response.Response{} "Successfully deleted {{.EntityPascal}}"
// @Router /{{.EntityKebab}}/{id} [delete]
func (con  *{{.EntityPascal}}Controller) Delete(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Delete("/:id").Name("delete one {{.EntityPascal}}").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID int %sparam:"id" validate:"gt=0"%s
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				if err := con.service.Delete{{.EntityPascal}}(c.Request().Context(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The {{.EntityAllLower}} retrieved successfully!"),
				)
			},
		}
	})
}
`,
	"%s", special,
)
