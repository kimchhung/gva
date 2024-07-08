package menu

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"github.com/gva/api/admin/module/menu/dto"

	permissions "github.com/gva/app/common/permission"
	"github.com/gva/app/common/service"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/echoc"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/internal/rql"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*RouteController)(nil)

type RouteController struct {
	menu_s *MenuService
	jwt_s  *service.JwtService
	log    *zerolog.Logger
}

func NewMenuController(
	menu_s *MenuService,
	jwt_s *service.JwtService,
	log *zerolog.Logger,
) *RouteController {
	return &RouteController{
		menu_s: menu_s,
		jwt_s:  jwt_s,
		log:    log,
	}
}

func (con *RouteController) Init(r *echo.Group) *echo.Group {
	return r.Group("/menus", con.jwt_s.RequiredAdmin())
}

// @Tags        Route
// @Summary     List all Routes
// @Description Get a list of all Routes
// @ID          list-all-routes
// @Produce     json
// @Success     200 {object} response.Response{data=map[string]dto.MenuResponse{list=[]dto.MenuResponse}} "Successfully retrieved Routes"
// @Router      /routes [get]
// @Security    Bearer
// @Param   	limit     query     int     false  "string default"     default(A)
func (con *RouteController) Routes(m *echoc.RouteMeta) echoc.MetaHandler {
	parser := request.MustRqlParser(rql.Config{
		// Table:        route.Table,
		Model:        ent.Menu{},
		DefaultLimit: 20,
		DefaultSort:  []string{"-id"},
		FieldSep:     ".",
	})

	return m.Get("/").DoWithScope(func() []echo.HandlerFunc {
		params := new(dto.MenuPaginateRequest)
		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.menu_s.Paginate(c.Request().Context(), params)
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

// @Tags        Route
// @Summary     Create a Route
// @Description Create a Route
// @ID          create-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.MenuResponse} "Successfully created Routes"
// @Router      /routes [post]
// @Security    Bearer
// @Param 		info body dto.MenuRequest true "Route Info"
func (con *RouteController) CreateRoute(m *echoc.RouteMeta) echoc.MetaHandler {
	return m.Post("/").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.MenuRequest)

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(request.BodyParser(body)),
			func(c echo.Context) error {
				data, err := con.menu_s.CreateRoute(c.Request().Context(), body)
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

// @Tags        Route
// @Summary     Get a Route
// @Description Get a Route
// @ID          Get-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.MenuResponse} "Successfully Getd Routes"
// @Router      /routes/{id} [put]
// @Security    Bearer
// @Param 		info body dto.MenuRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) GetRoute(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Put("/:id").DoWithScope(func() []echo.HandlerFunc {
		params := new(struct {
			ID xid.ID `param:"id" validate:"required"`
		})

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.menu_s.GetRouteByID(c.Request().Context(), params.ID)
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

// @Tags        Route
// @Summary     Update a Route
// @Description Update a Route
// @ID          Update-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.MenuResponse} "Successfully updated Routes"
// @Router      /routes/{id} [put]
// @Security    Bearer
// @Param 		info body dto.MenuRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) UpdateRoute(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Put("/:id").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.MenuRequest)
		params := new(struct {
			ID xid.ID `param:"id" validate:"required"`
		})

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.BodyParser(body),
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.menu_s.UpdateRoute(c.Request().Context(), params.ID, body)
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

// @Tags        Route
// @Summary     Delete a Route
// @Description Delete a Route
// @ID          Delete-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{} "Successfully Delete Routes"
// @Router      /routes/{id} [delete]
// @Security    Bearer
// @Param 		id path int true "Route ID"
func (con *RouteController) DeleteRoute(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Delete("/:id").DoWithScope(func() []echo.HandlerFunc {
		params := new(struct {
			ID xid.ID `param:"id" validate:"required"`
		})

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.menu_s.DeleteRoute(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c)
			},
		}
	})
}
