# Generate Resource Components

This document provides commands to generate various components like Resource Controller, Service, Repository, and Model in your project.

## Generate Resource

- `make admincmd.gen name="todo"`: Generate all components in a resource.
- `make admincmd.gen name="todo" option=MD`: Generate only model in a resource.
- `make admincmd.gen name="todo" option=P`: Generate only permission in a resource.
- `make admincmd.gen name="todo" option=M`: Generate only module in a resource.
- `make admincmd.gen name="todo" option=C`: Generate only controller in a resource.
- `make admincmd.gen name="todo" option=S`: Generate only service in a resource.
- `make admincmd.gen name="todo" option=M,C,S,D`: Generate module, controller, service and dto in a resource.

## Generate Swagger `make swag` will generate swagger docs base on this decoration

```
// @Tags			Admin
// @Summary			Get Admin
// @Description		Get Admin
// @ID				get-admin
// @Accept			json
// @Param			id	path	int	true	"Admin ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.AdminResponse}	"Successfully get Admin"
// @Router			/admin/{id} [get]
func (con *AdminController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetAdminRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.AdminView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetAdmin(c.Request().Context(), params.ID)
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
```
