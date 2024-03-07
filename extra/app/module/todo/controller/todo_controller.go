package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/app/module/todo/dto"
	"github.com/kimchhung/gva/extra/app/module/todo/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*TodoController)(nil)

type TodoController struct {
	service *service.TodoService
	jwtSv   *services.JwtService
}

func (con *TodoController) Routes(r fiber.Router) {
	todo := r.Group("todo")
	todo.Use(con.jwtSv.ProtectAdmin())
	rctrl.Register(todo, con)
}

func NewTodoController(service *service.TodoService, jwtSv *services.JwtService) *TodoController {
	return &TodoController{
		service: service,
		jwtSv:   jwtSv,
	}
}

// @Tags Todo
// @Summary List all Todos
// @Description Get a list of all Todos
// @ID list-all-Todos
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.TodoResponse{list=[]dto.TodoResponse}} "Successfully retrieved Todos"
// @Router /todo [get]
// @Security Bearer
func (con *TodoController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Todos").Do(func(c *fiber.Ctx) error {
		list, err := con.service.GetTodos(c.UserContext())
		if err != nil {
			return err
		}

		return request.Response(c,
			response.Data(list),
			response.Message("Todo list retreived successfully!"),
		)
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
func (con *TodoController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Todo").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.GetTodoByID(c.UserContext(), param.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The todo retrieved successfully!"),
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
// @Success  200 {object} response.Response{data=dto.TodoResponse} "Successfully created Todo"
// @Router /todo [post]
func (con *TodoController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Todo").DoWithScope(func() []fiber.Handler {
		body := new(dto.TodoRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreateTodo(c.UserContext(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The todo retrieved successfully!"),
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
// @Success  200 {object} response.Response{data=dto.TodoResponse} "Successfully updated Todo"
// @Router /todo/{id} [patch]
func (con *TodoController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Todo").DoWithScope(func() []fiber.Handler {
		body := new(dto.TodoRequest)
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.UpdateTodo(c.UserContext(), param.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The todo retrieved successfully!"),
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
// @Success  200 {object} response.Response{} "Successfully deleted Todo"
// @Router /todo/{id} [delete]
func (con *TodoController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Todo").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.DeleteTodo(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The todo retrieved successfully!"),
				)
			},
		}
	})
}
