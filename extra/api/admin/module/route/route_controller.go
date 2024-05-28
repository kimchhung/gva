package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/kimchhung/gva/extra/api/admin/module/route/dto"

	permissions "github.com/kimchhung/gva/extra/app/common/permission"
	"github.com/kimchhung/gva/extra/app/common/service"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/internal/rql"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*RouteController)(nil)

type RouteController struct {
	route_s *RouteService
	jwt_s   *service.JwtService
	log     *zerolog.Logger
}

func (con *RouteController) Init(r fiber.Router) fiber.Router {
	return r.Group("route")
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

// @Tags        Route
// @Summary     List all Routes
// @Description Get a list of all Routes
// @ID          list-all-routes
// @Produce     json
// @Success     200 {object} response.Response{data=map[string]dto.RouteResponse{list=[]dto.RouteResponse}} "Successfully retrieved Routes"
// @Router      /route [get]
// @Security    Bearer
// @Param   	limit     query     int     false  "string default"     default(A)
func (con *RouteController) Routes(m *rctrl.RouteMeta) rctrl.MetaHandler {
	parser := request.MustRqlParser(rql.Config{
		// Table:        route.Table,
		Model:        ent.Route{},
		DefaultLimit: 20,
		DefaultSort:  []string{"-id"},
		FieldSep:     ".",
	})

	return m.Get("/").DoWithScope(func() []fiber.Handler {
		params := new(dto.RoutePaginateRequest)

		return []fiber.Handler{
			request.Parse(
				request.RqlQueryParser(&params.Params, parser),
				request.QueryParser(params),
			),
			func(c *fiber.Ctx) error {
				list, meta, err := con.route_s.Paginate(c.UserContext(), params)
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
// @Router      /route [post]
// @Security    Bearer
// @Param 		info body dto.RouteRequest true "Route Info"
func (con *RouteController) CreateRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").DoWithScope(func() []fiber.Handler {
		body := new(dto.RouteRequest)

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Validate(request.BodyParser(body)),

			func(c *fiber.Ctx) error {
				data, err := con.route_s.CreateRoute(c.UserContext(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Meta(meta),
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
// @Router      /route [put]
// @Security    Bearer
// @Param 		info body dto.RouteRequest true "Route Info"
// @Param 		id path int true "Route ID"
func (con *RouteController) UpdateRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Put("/:id").DoWithScope(func() []fiber.Handler {
		body := new(dto.RouteRequest)
		params := new(struct {
			ID int `params:"id" validate:"required,min=0"`
		})

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Validate(
				request.BodyParser(body),
				request.ParamsParser(params),
			),
			func(c *fiber.Ctx) error {
				data, err := con.route_s.UpdateRoute(c.UserContext(), params.ID, body)
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
// @Router      /route [delete]
// @Security    Bearer
// @Param 		id path int true "Route ID"
func (con *RouteController) DeleteRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").DoWithScope(func() []fiber.Handler {
		params := new(struct {
			ID int `params:"id" validate:"required,min=0"`
		})

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c *fiber.Ctx) error {
				err := con.route_s.DeleteRoute(c.UserContext(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c)
			},
		}
	})
}
