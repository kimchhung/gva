package todo

import (
	"backend/app/admin/middleware"
	"backend/app/admin/module/todo/dto"
	"backend/app/share/permission"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*TodoController)(nil)

type TodoController struct {
	middleware *middleware.Middleware
	service    *TodoService
}

func NewTodoController(
	middleware *middleware.Middleware,
	service *TodoService,
) *TodoController {
	return &TodoController{
		middleware: middleware,
		service:    service,
	}
}

func (con *TodoController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/todo"),
	)
}

// @Tags			Todo
// @Summary			Create Todo
// @Description		Create Todo
// @ID				create-todo
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.CreateTodoRequest	true	"Todo data"
// @Success			200		{object}	response.Response{data=dto.TodoResponse}	"Successfully created Todo"
// @Router			/todo [post]
func (con *TodoController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.CreateTodoRequest)

		return []ctr.H{
			permission.RequireAny(permission.TodoAdd),
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

// @Tags			Todo
// @Summary			Get Todo
// @Description		Get Todo
// @ID				get-todo
// @Accept			json
// @Param			id	path	int	true	"Todo ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.TodoResponse}	"Successfully get Todo"
// @Router			/todo/{id} [get]
func (con *TodoController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetTodoRequest)

		return []ctr.H{
			permission.RequireAny(permission.TodoView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetTodo(c.Request().Context(), params.ID)
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

// @Tags			Todo
// @Summary			Get Todos
// @Description		Get Todos
// @ID				get-todos
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
// @Success			200		{object}	response.Response{data=[]dto.TodoResponse}	"Successfully get Todos"
// @Router			/todo [get]
func (con *TodoController) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAny(permission.TodoView),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.GetTodos(c.Request().Context(), query)
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

// @Tags			Todo
// @Summary			Update Todo
// @Description		Update Todo
// @ID				update-todo
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Todo ID"
// @Param			body	body		dto.UpdateTodoRequest	true	"Todo data"
// @Success			200		{object}	response.Response{data=dto.TodoResponse}	"Successfully updated Todo"
// @Router			/todo/{id} [put]
func (con *TodoController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.GetTodoRequest)
		body := new(dto.UpdateTodoRequest)

		return []ctr.H{
			permission.RequireAny(permission.TodoEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateTodo(c.Request().Context(), params.ID, body)
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

// @Tags			Todo
// @Summary			Update Todo partial
// @Description		Update Todo partial
// @ID				update-todo-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "Todo ID"
// @Param			body	body	dto.UpdatePatchTodoRequest	true	"Todo data"
// @Success			200		{object}	response.Response{data=dto.TodoResponse}	"Successfully updated Todo"
// @Router			/todo/{id} [patch]
func (con *TodoController) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.GetTodoRequest)
		body := new(dto.UpdatePatchTodoRequest)

		return []ctr.H{
			permission.RequireAny(permission.TodoEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePatchTodo(c.Request().Context(), params.ID, body)
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

// @Tags			Todo
// @Summary			Delete Todo
// @Description		Delete Todo
// @ID				delete-todo
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Todo ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted Todo"
// @Router			/todo/{id} [delete]
func (con *TodoController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.GetTodoRequest)

		return []ctr.H{
			permission.RequireAny(permission.TodoDelete),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeleteTodo(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted Todo"),
				)
			},
		}
	})
}
