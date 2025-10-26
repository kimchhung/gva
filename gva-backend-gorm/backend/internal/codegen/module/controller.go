package module_template

import (
	"strings"
)

// {{.EntityAllLower}}
// {{.EntityPascal}}

var Controller = strings.ReplaceAll(`package {{.EntityAllLower}}

import (
	"backend/app/admin/middleware"
	"backend/app/admin/module/{{.EntityAllLower}}/dto"
	"backend/app/share/permission"
	"backend/internal/ctr"
	"backend/core/utils/request"
	"backend/core/utils/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*{{.EntityPascal}}Controller)(nil)

type {{.EntityPascal}}Controller struct {
	middleware *middleware.Middleware
	service *{{.EntityPascal}}Service
}

func New{{.EntityPascal}}Controller(
	middleware *middleware.Middleware,
	service *{{.EntityPascal}}Service,
) *{{.EntityPascal}}Controller {
	return &{{.EntityPascal}}Controller{
		middleware: middleware,
		service: service,
	}
}

func (con *{{.EntityPascal}}Controller) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/{{.EntityKebab}}"),
	)
}

// @Tags			{{.EntityPascal}}
// @Summary			Create {{.EntityPascal}}
// @Description		Create {{.EntityPascal}}
// @ID				create-{{.EntityAllLower}}
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.Create{{.EntityPascal}}Request	true	"{{.EntityPascal}} data"
// @Success			200		{object}	response.Response{data=dto.{{.EntityPascal}}Response}	"Successfully created {{.EntityPascal}}"
// @Router			/{{.EntityKebab}} [post]
func (con *{{.EntityPascal}}Controller) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.Create{{.EntityPascal}}Request)

		return []ctr.H{
			permission.RequireAny(permission.{{.EntityPascal}}Add),
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.Create{{.EntityPascal}}(c.Request().Context(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			{{.EntityPascal}}
// @Summary			Get {{.EntityPascal}}
// @Description		Get {{.EntityPascal}}
// @ID				get-{{.EntityKebab}}
// @Accept			json
// @Param			id	path	int	true	"{{.EntityPascal}} ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.{{.EntityPascal}}Response}	"Successfully get {{.EntityPascal}}"
// @Router			/{{.EntityKebab}}/{id} [get]
func (con *{{.EntityPascal}}Controller) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.Get{{.EntityPascal}}Request)

		return []ctr.H{
			permission.RequireAny(permission.{{.EntityPascal}}View),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.Get{{.EntityPascal}}(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			{{.EntityPascal}}
// @Summary			Get {{.EntityPascal}}s
// @Description		Get {{.EntityPascal}}s
// @ID				get-{{.EntityKebab}}s
// @Accept			json
// @Param 			page	query	int	false	"page"
// @Param 			limit	query	int	false	"limit"
// @Param 			search	query	string	false	"search"
// @Param 			filters	query	string	false	"filters"
// @Param 			sorts	query	string	false	"sorts"
// @Param 			orders	query	string	false	"orders"
// @Param 			selects	query	string	false	"selects: list, totalCount"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=[]dto.{{.EntityPascal}}Response}	"Successfully get {{.EntityPascal}}s"
// @Router			/{{.EntityKebab}} [get]
func (con *{{.EntityPascal}}Controller) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAny(permission.{{.EntityPascal}}View),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.Get{{.EntityPascal}}s(c.Request().Context(), query)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(list),
					response.Meta(meta),
				)
			},
		}
	})
}

// @Tags			{{.EntityPascal}}
// @Summary			Update {{.EntityPascal}}
// @Description		Update {{.EntityPascal}}
// @ID				update-{{.EntityAllLower}}
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"{{.EntityPascal}} ID"
// @Param			body	body		dto.Update{{.EntityPascal}}Request	true	"{{.EntityPascal}} data"
// @Success			200		{object}	response.Response{data=dto.{{.EntityPascal}}Response}	"Successfully updated {{.EntityPascal}}"
// @Router			/{{.EntityKebab}}/{id} [put]
func (con *{{.EntityPascal}}Controller) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.Get{{.EntityPascal}}Request)
		body := new(dto.Update{{.EntityPascal}}Request)

		return []ctr.H{
			permission.RequireAny(permission.{{.EntityPascal}}Edit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.Update{{.EntityPascal}}(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			{{.EntityPascal}}
// @Summary			Update {{.EntityPascal}} partial
// @Description		Update {{.EntityPascal}} partial
// @ID				update-{{.EntityAllLower}}-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "{{.EntityPascal}} ID"
// @Param			body	body	dto.UpdatePatch{{.EntityPascal}}Request	true	"{{.EntityPascal}} data"
// @Success			200		{object}	response.Response{data=dto.{{.EntityPascal}}Response}	"Successfully updated {{.EntityPascal}}"
// @Router			/{{.EntityKebab}}/{id} [patch]
func (con *{{.EntityPascal}}Controller) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.Get{{.EntityPascal}}Request)
		body := new(dto.UpdatePatch{{.EntityPascal}}Request)

		return []ctr.H{
			permission.RequireAny(permission.{{.EntityPascal}}Edit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePatch{{.EntityPascal}}(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			{{.EntityPascal}}
// @Summary			Delete {{.EntityPascal}}
// @Description		Delete {{.EntityPascal}}
// @ID				delete-{{.EntityAllLower}}
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"{{.EntityPascal}} ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted {{.EntityPascal}}"
// @Router			/{{.EntityKebab}}/{id} [delete]
func (con *{{.EntityPascal}}Controller) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.Get{{.EntityPascal}}Request)

		return []ctr.H{
			permission.RequireAny(permission.{{.EntityPascal}}Delete),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.Delete{{.EntityPascal}}(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted {{.EntityPascal}}"),
				)
			},
		}
	})
}
`,
	"%s", special,
)
