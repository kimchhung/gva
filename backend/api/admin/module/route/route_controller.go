package route

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"github.com/kimchhung/gva/backend/api/admin/module/route/dto"

	permissions "github.com/kimchhung/gva/backend/app/common/permission"
	"github.com/kimchhung/gva/backend/app/common/service"
	"github.com/kimchhung/gva/backend/internal/echoc"
	"github.com/kimchhung/gva/backend/internal/ent"
	"github.com/kimchhung/gva/backend/internal/request"
	"github.com/kimchhung/gva/backend/internal/response"
	"github.com/kimchhung/gva/backend/internal/rql"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*RouteController)(nil)

type RouteController struct {
	route_s *RouteService
	jwt_s   *service.JwtService
	log     *zerolog.Logger
}

func NewRouteController(
	route_s *RouteService,
	jwt_s *service.JwtService,
	log *zerolog.Logger,
) *RouteController {
	return &RouteController{
		route_s: route_s,
		jwt_s:   jwt_s,
		log:     log,
	}
}

func (con *RouteController) Init(r *echo.Group) *echo.Group {
	return r.Group("/routes")
}

// @Tags        Route
// @Summary     List all Routes
// @Description Get a list of all Routes
// @ID          list-all-routes
// @Produce     json
// @Success     200 {object} response.Response{data=map[string]dto.RouteResponse{list=[]dto.RouteResponse}} "Successfully retrieved Routes"
// @Router      /routes [get]
// @Security    Bearer
// @Param   	limit     query     int     false  "string default"     default(A)
func (con *RouteController) Routes(m *echoc.RouteMeta) echoc.MetaHandler {
	parser := request.MustRqlParser(rql.Config{
		// Table:        route.Table,
		Model:        ent.Route{},
		DefaultLimit: 20,
		DefaultSort:  []string{"-id"},
		FieldSep:     ".",
	})

	return m.Get("/").DoWithScope(func() []echo.HandlerFunc {
		params := new(dto.RoutePaginateRequest)
		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c echo.Context) error {
				list, meta, err := con.route_s.Paginate(c.Request().Context(), params)
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
// @Success     200 {object} response.Response{data=dto.RouteResponse} "Successfully created Routes"
// @Router      /routes [post]
// @Security    Bearer
// @Param 		info body dto.RouteRequest true "Route Info"
func (con *RouteController) CreateRoute(m *echoc.RouteMeta) echoc.MetaHandler {
	return m.Post("/").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.RouteRequest)

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(request.BodyParser(body)),
			func(c echo.Context) error {
				data, err := con.route_s.CreateRoute(c.Request().Context(), body)
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
// @Success     200 {object} response.Response{data=dto.RouteResponse} "Successfully Getd Routes"
// @Router      /routes/{id} [put]
// @Security    Bearer
// @Param 		info body dto.RouteRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) GetRoute(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Put("/:id").DoWithScope(func() []echo.HandlerFunc {
		params := new(struct {
			ID int `param:"id" validate:"required,min=0"`
		})

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.route_s.GetRouteByID(c.Request().Context(), params.ID)
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
// @Success     200 {object} response.Response{data=dto.RouteResponse} "Successfully updated Routes"
// @Router      /routes/{id} [put]
// @Security    Bearer
// @Param 		info body dto.RouteRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) UpdateRoute(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Put("/:id").DoWithScope(func() []echo.HandlerFunc {
		body := new(dto.RouteRequest)
		params := new(struct {
			ID int `param:"id" validate:"required,min=0"`
		})

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.BodyParser(body),
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.route_s.UpdateRoute(c.Request().Context(), params.ID, body)
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
			ID int `param:"id" validate:"required,min=0"`
		})

		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.route_s.DeleteRoute(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c)
			},
		}
	})
}
