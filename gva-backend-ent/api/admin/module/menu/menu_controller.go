package menu

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"github.com/gva/api/admin/module/menu/dto"

	permissions "github.com/gva/app/common/permission"
	"github.com/gva/app/common/service"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/internal/rql"
)

// don't remove for runtime type checking
var _ interface{ ctr.CTR } = (*RouteController)(nil)

type RouteController struct {
	service *MenuService
	jwt_s   *service.JwtService
	log     *zerolog.Logger
}

func NewMenuController(
	service *MenuService,
	jwt_s *service.JwtService,
	log *zerolog.Logger,
) *RouteController {
	return &RouteController{
		service: service,
		jwt_s:   jwt_s,
		log:     log,
	}
}

func (con *RouteController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/menu", con.jwt_s.RequiredAdmin()),
	)
}

// @Tags        Menu
// @Summary     List all Menus
// @Description Get a list of all Menus
// @ID          list-all-Menus
// @Produce     json
// @Success     200 {object} response.Response{data=map[string]dto.MenuResponse{list=[]dto.MenuResponse}}
// @Router      /menu [get]
// @Security    Bearer
// @Param   	limit     query     int     false  "string default"     default(A)
func (con *RouteController) List() *ctr.Route {
	parser := request.MustRqlParser(rql.Config{
		Model: struct {
			ID pxid.ID `json:"id" rql:"filter,sort"`
		}{},
	})

	return ctr.GET("/").Do(func() []ctr.H {
		var (
			params = new(dto.MenuPagedRequest)
		)
		return []ctr.H{
			permissions.OnlySuperAdmin(),
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.service.Paginate(c.Request().Context(), params)
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

// @Tags        Menu
// @Summary     List all Menus
// @Description Get a list of all Enabled Menus
// @ID          list-all-public-menus
// @Produce     json
// @Success     200 {object} response.Response{data=map[string]dto.MenuResponse{list=[]dto.MenuResponse}}
// @Router      /menu/enabled-list [get]
// @Security    Bearer
// @Param   	limit     query     int     false  "string default"     default(A)
func (con *RouteController) EnableList() *ctr.Route {
	return ctr.GET("/enabled-list").Do(func() []ctr.H {
		return []ctr.H{func(c echo.Context) error {
			list, err := con.service.EnabledList(c.Request().Context())
			if err != nil {
				return err
			}

			return request.Response(c,
				response.Data(list),
			)
		}}
	})
}

// @Tags        Menu
// @Summary     Get a Route
// @Description Get a Route
// @ID          Get-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.MenuResponse} "Successfully Getd Routes"
// @Router      /menu/{id} [put]
// @Security    Bearer
// @Param 		info body dto.MenuRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(struct {
			ID pxid.ID `param:"id" validate:"required"`
		})

		return []ctr.H{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetMenuByID(c.Request().Context(), params.ID)
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

// @Tags        Menu
// @Summary     Create a Menu
// @Description Create a Menu
// @ID          create-a-menu
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.MenuResponse}
// @Router      /menu [post]
// @Security    Bearer
// @Param 		info body dto.MenuRequest true "Route Info"
func (con *RouteController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.MenuRequest)

		return []ctr.H{
			permissions.OnlySuperAdmin(),
			request.Validate(request.BodyParser(body)),
			func(c echo.Context) error {
				data, err := con.service.CreateMenu(c.Request().Context(), body)
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

// @Tags        Menu
// @Summary     Update a Menu
// @Description Update a Menu
// @ID          Update-a-Menu
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.MenuResponse}
// @Router      /menu/{id} [put]
// @Security    Bearer
// @Param 		info body dto.MenuRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		body := new(dto.MenuRequest)
		params := new(struct {
			ID pxid.ID `param:"id" validate:"required"`
		})

		return []ctr.H{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.BodyParser(body),
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateMenu(c.Request().Context(), params.ID, body)
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

// @Tags        Menu
// @Summary     Delete a Menu
// @Description Delete a Menu
// @ID          Delete-a-Menu
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{}
// @Router      /menu/{id} [delete]
// @Security    Bearer
// @Param 		id path int true "Route ID"
func (con *RouteController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(struct {
			ID pxid.ID `param:"id" validate:"required"`
		})

		return []ctr.H{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeleteMenu(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c)
			},
		}
	})
}
