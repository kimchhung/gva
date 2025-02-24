package department

import (
	"github.com/gva/api/admin/module/department/dto"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/internal/rql"
	"github.com/labstack/echo/v4"
)

// don't remove for runtime type checking
var _ interface{ ctr.CTR } = (*DepartmentController)(nil)

type DepartmentController struct {
	service *DepartmentService
}

func (con *DepartmentController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/department"),
	)
}

func NewDepartmentController(service *DepartmentService) *DepartmentController {
	return &DepartmentController{
		service: service,
	}
}

// @Tags Department
// @Summary List all Departments
// @Description Get a list of all Departments
// @ID list-all-Departments
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.DepartmentResponse{list=[]dto.DepartmentResponse}} "Successfully retrieved Departments"
// @Router /department [get]
// @Security Bearer
func (con *DepartmentController) List() *ctr.Route {
	parser := request.MustRqlParser(rql.Config{
		Model: struct {
			ID pxid.ID `json:"id" rql:"filter,sort"`
		}{},
	})

	return ctr.GET("/").Do(func() []ctr.H {
		params := new(dto.DepartmentPagedRequest)
		return []ctr.H{
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.service.GetDepartments(c.Request().Context(), params)
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

// @Tags Department
// @Security Bearer
// @Summary Get a Department
// @Description Get a Department by ID
// @ID get-Department-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Department ID"
// @Success   200 {object} response.Response{data=dto.DepartmentResponse}
// @Router /department/{id} [get]
func (con *DepartmentController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		param := &struct {
			ID pxid.ID `param:"id" validate:"required"`
		}{}

		return []ctr.H{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				data, err := con.service.GetDepartmentByID(c.Request().Context(), param.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The department retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Department
// @Security Bearer
// @Summary Create a Department
// @Description Create a new Department with the provided details
// @ID create-Department
// @Accept  json
// @Produce  json
// @Param Department body dto.DepartmentRequest true "Department data"
// @Success  200 {object} response.Response{data=dto.DepartmentResponse} "Successfully created Department"
// @Router /department [post]
func (con *DepartmentController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.DepartmentRequest)

		return []ctr.H{
			request.Validate(
				request.BodyParser(body),
			),

			func(c echo.Context) error {
				data, err := con.service.CreateDepartment(c.Request().Context(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The department created successfully!"),
				)
			},
		}
	})
}

// @Tags Department
// @Security Bearer
// @Summary Update a Department
// @Description Update a Department by ID
// @ID update-Department-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Department ID"
// @Param Department body dto.DepartmentRequest true "Department data"
// @Success  200 {object} response.Response{data=dto.DepartmentResponse} "Successfully updated Department"
// @Router /department/{id} [patch]
func (con *DepartmentController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		body := new(dto.DepartmentRequest)
		param := &struct {
			ID pxid.ID `param:"id" validate:"required"`
		}{}

		return []ctr.H{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateDepartment(c.Request().Context(), param.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The department updated successfully!"),
				)
			},
		}
	})
}

// @Tags Department
// @Security Bearer
// @Summary Delete a Department
// @Description Delete a Department by ID
// @ID delete-Department-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Department ID"
// @Success  200 {object} response.Response{} "Successfully deleted Department"
// @Router /department/{id} [delete]
func (con *DepartmentController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Name("delete one Department").Do(func() []ctr.H {
		param := &struct {
			ID pxid.ID `param:"id" validate:"required"`
		}{}

		return []ctr.H{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c echo.Context) error {
				if err := con.service.DeleteDepartment(c.Request().Context(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The department deleted successfully!"),
				)
			},
		}
	})
}
