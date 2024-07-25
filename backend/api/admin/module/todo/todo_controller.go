package todo

import (
	"github.com/gva/api/admin/module/todo/dto"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/echoc"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/internal/rql"
	"github.com/labstack/echo/v4"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*TodoController)(nil)

type TodoController struct {
	service *TodoService
}

func (con *TodoController) Init(r *echo.Group) *echo.Group {
	return r.Group("/todo")
}

func NewTodoController(service *TodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}

// @Tags Todo
// @Summary List all Todos
// @Description Get a list of all Todos
// @ID list-all-Todos
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.TodoResponse{list=[]dto.TodoResponse}}"
// @Router /todo [get]
// @Security Bearer
func (con *TodoController) List(meta *echoc.RouteMeta) echoc.MetaHandler {
	parser := request.MustRqlParser(rql.Config{
		Model: struct {
			ID xid.ID `json:"id" rql:"filter,sort"`
		}{},
	})
	return meta.Get("/").DoWithScope(func() []echo.HandlerFunc {
		params := new(dto.TodoPagedRequest)
		return []echo.HandlerFunc{
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.service.GetTodos(c.Request().Context(), params)
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

// @Tags Todo
// @Security Bearer
// @Summary Get a Todo
// @Description Get a Todo by ID
// @ID get-Todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success   200 {object} response.Response{data=dto.TodoResponse}
// @Router /todo/{id} [get]
func (con *TodoController) Get(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/:id").Name("get one Todo").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				data, err := con.service.GetTodoByID(c.Request().Context(), param.ID)
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

// @Tags Todo
// @Security Bearer
// @Summary Create a Todo
// @Description Create a new Todo with the provided details
// @ID create-Todo
// @Accept  json
// @Produce  json
// @Param Todo body dto.TodoRequest true "Todo data"
// @Success  200 {object} response.Response{data=dto.TodoResponse}
// @Router /todo [post]
func (con *TodoController) Create(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Post("/").Name("create one Todo").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.TodoRequest)

		return []echo.HandlerFunc{
			request.Validate(
				request.BodyParser(body),
			),

			func(c echo.Context) error {
				data, err := con.service.CreateTodo(c.Request().Context(), body)
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

// @Tags Todo
// @Security Bearer
// @Summary Update a Todo
// @Description Update a Todo by ID
// @ID update-Todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param Todo body dto.TodoRequest true "Todo data"
// @Success  200 {object} response.Response{data=dto.TodoResponse}
// @Router /todo/{id} [patch]
func (con *TodoController) Update(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Patch("/:id").Name("update one Todo").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.TodoRequest)
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateTodo(c.Request().Context(), param.ID, body)
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

// @Tags Todo
// @Security Bearer
// @Summary Delete a Todo
// @Description Delete a Todo by ID
// @ID delete-Todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success  200 {object} response.Response{} "The todo deleted successfully!"
// @Router /todo/{id} [delete]
func (con *TodoController) Delete(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Delete("/:id").Name("delete one Todo").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				if err := con.service.DeleteTodo(c.Request().Context(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The todo deleted successfully!"),
				)
			},
		}
	})
}
