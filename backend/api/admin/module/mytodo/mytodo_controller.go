package mytodo

import (
	"github.com/gva/api/admin/module/mytodo/dto"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/echoc"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/labstack/echo/v4"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*MyTodoController)(nil)

type MyTodoController struct {
	service *MyTodoService
}

func (con *MyTodoController) Init(r *echo.Group) *echo.Group {
	return r.Group("/my-todo")
}

func NewMyTodoController(service *MyTodoService) *MyTodoController {
	return &MyTodoController{
		service: service,
	}
}

// @Tags MyTodo
// @Summary List all MyTodos
// @Description Get a list of all MyTodos
// @ID list-all-MyTodos
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.MyTodoResponse{list=[]dto.MyTodoResponse}} "Successfully retrieved MyTodos"
// @Router /my-todo [get]
// @Security Bearer
func (con *MyTodoController) List(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/").Name("get many MyTodos").Do(func(c echo.Context) error {
		list, err := con.service.GetMyTodos(c.Request().Context())
		if err != nil {
			return err
		}

		return request.Response(c,
			response.Data(list),
			response.Message("MyTodo list retreived successfully!"),
		)
	})
}

// @Tags MyTodo
// @Security Bearer
// @Summary Get a MyTodo
// @Description Get a MyTodo by ID
// @ID get-MyTodo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "MyTodo ID"
// @Success   200 {object} response.Response{data=dto.MyTodoResponse}
// @Router /my-todo/{id} [get]
func (con *MyTodoController) Get(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/:id").Name("get one MyTodo").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				data, err := con.service.GetMyTodoByID(c.Request().Context(), param.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The mytodo retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags MyTodo
// @Security Bearer
// @Summary Create a MyTodo
// @Description Create a new MyTodo with the provided details
// @ID create-MyTodo
// @Accept  json
// @Produce  json
// @Param MyTodo body dto.MyTodoRequest true "MyTodo data"
// @Success  200 {object} response.Response{data=dto.MyTodoResponse} "Successfully created MyTodo"
// @Router /my-todo [post]
func (con *MyTodoController) Create(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Post("/").Name("create one MyTodo").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.MyTodoRequest)

		return []echo.HandlerFunc{
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.CreateMyTodo(c.Request().Context(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The mytodo retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags MyTodo
// @Security Bearer
// @Summary Update a MyTodo
// @Description Update a MyTodo by ID
// @ID update-MyTodo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "MyTodo ID"
// @Param MyTodo body dto.MyTodoRequest true "MyTodo data"
// @Success  200 {object} response.Response{data=dto.MyTodoResponse} "Successfully updated MyTodo"
// @Router /my-todo/{id} [patch]
func (con *MyTodoController) Update(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Patch("/:id").Name("update one MyTodo").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.MyTodoRequest)
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateMyTodo(c.Request().Context(), param.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The mytodo retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags MyTodo
// @Security Bearer
// @Summary Delete a MyTodo
// @Description Delete a MyTodo by ID
// @ID delete-MyTodo-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "MyTodo ID"
// @Success  200 {object} response.Response{} "Successfully deleted MyTodo"
// @Router /my-todo/{id} [delete]
func (con *MyTodoController) Delete(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Delete("/:id").Name("delete one MyTodo").DoWithScope(func() []echo.HandlerFunc {
		param := &struct {
			ID xid.ID `param:"id" validate:"required"`
		}{}

		return []echo.HandlerFunc{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				if err := con.service.DeleteMyTodo(c.Request().Context(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The mytodo retrieved successfully!"),
				)
			},
		}
	})
}
