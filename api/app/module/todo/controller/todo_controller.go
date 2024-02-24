package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/module/todo/dto"
	"github.com/kimchhung/gva/app/module/todo/service"
	"github.com/kimchhung/gva/internal/rctrl"
	"github.com/kimchhung/gva/utils/request"
)

// don't remove for runtime type checking
var _ ITodoController = (*TodoController)(nil)

type ITodoController interface {
	rctrl.FiberRouter
	Create(meta *rctrl.RouteMeta) rctrl.MetaHandler
	List(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Get(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Update(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler
}


type TodoController struct {
	service *service.TodoService
}

func (con *TodoController) Routes(r fiber.Router) {
	r.Route("todo",
		func(router fiber.Router) {
			rctrl.Register(router, con)
		},
	)
}

func NewTodoController(service *service.TodoService) *TodoController {
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
// @Success  200 {object} request.Response{data=[]dto.TodoResponse} "Successfully retrieved Todos"
// @Router /todo [get]
func (con *TodoController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Todos").Do(func(c *fiber.Ctx) error {
		list, err := con.service.GetTodos(c.UserContext())
		if err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "Todo list retreived successfully!",
			Data:    list,
		})
	})
}

// @Tags Todo
// @Summary Get a Todo
// @Description Get a Todo by ID
// @ID get-Todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success   200 {object} request.Response{data=dto.TodoResponse}
// @Router /Todo/{id} [get]
func (con *TodoController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Todo").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
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

				return request.Resp(c, request.Response{
					Message: "The todo retrieved successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Todo
// @Summary Create a Todo
// @Description Create a new Todo with the provided details
// @ID create-Todo
// @Accept  json
// @Produce  json
// @Param Todo body dto.TodoRequest true "Todo data"
// @Success  200 {object} request.Response{data=dto.TodoResponse} "Successfully created Todo"
// @Router /todo [post]
func (con *TodoController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Todo").DoWithScope(func() []fiber.Handler {
		req := new(dto.TodoRequest)

		return []fiber.Handler{
			request.Validate(request.BodyParser(&req)),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreateTodo(c.UserContext(), *req)
				if err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The todo was created successfully!",
					Data:    data,
				})
			},
		}
	})
}


// @Tags Todo
// @Summary Update a Todo
// @Description Update a Todo by ID
// @ID update-Todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param Todo body dto.TodoRequest true "Todo data"
// @Success  200 {object} request.Response{data=dto.TodoResponse} "Successfully updated Todo"
// @Router /Todo/{id} [patch]
func (con *TodoController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Todo").DoWithScope(func() []fiber.Handler {
		req := new(dto.TodoRequest)
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(req),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.UpdateTodo(c.UserContext(), param.ID, *req)
				if err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The todo was updated successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Todo
// @Summary Delete a Todo
// @Description Delete a Todo by ID
// @ID delete-Todo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success  200 {object} request.Response{} "Successfully deleted Todo"
// @Router /todo/{id} [delete]
func (con  TodoController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Todo").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gte=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.DeleteTodo(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The Todo was deleted successfully!",
				})
			},
		}
	})
}
