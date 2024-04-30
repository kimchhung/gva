package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/dto"
	"github.com/kimchhung/gva/extra/api/admin/module/authorization/service"
	"github.com/kimchhung/gva/extra/app/common/permissions"
	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/internal/rql"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*AuthorizationController)(nil)

type AuthorizationController struct {
	permission_ *service.PermissionService
	route_      *service.RouteService
}

func (con *AuthorizationController) Init(r fiber.Router) fiber.Router {
	return r.Group("authorization")
}

func NewAuthorizationController(permission_ *service.PermissionService, route_ *service.RouteService) *AuthorizationController {
	return &AuthorizationController{
		permission_: permission_,
		route_:      route_,
	}
}

// @Tags        Authorization Management
// @Summary     List all Routes
// @Description Get a list of all Routes
// @ID          list-all-routes
// @Produce     json
// @Success     200 {object} response.Response{data=map[string]dto.RouteResponse{list=[]dto.RouteResponse}} "Successfully retrieved Routes"
// @Router      /routes [get]
// @Security    Bearer
func (con *AuthorizationController) Routes(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	// init parser once and reused
	parser := request.MustRqlParser("routes",
		struct{ ent.Route }{},
	)

	return meta.Get("/routes").DoWithScope(func() []fiber.Handler {
		params := new(rql.Params)

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Parse(
				request.RqlParser(params, parser),
			),
			func(c *fiber.Ctx) error {
				list, meta, err := con.route_.Paginate(c.UserContext(), params)
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

// @Tags        Authorization Management
// @Summary     List all permissions
// @Description Get a list of all permissions
// @ID          list-all-permissions
// @Produce     json
// @Success     200 {object} response.Response{data=[]dto.PermissionResponse} "Successfully retrieved Routes"
// @Router      /permissions [get]
// @Security    Bearer
func (con *AuthorizationController) Permissions(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/permissions").DoWithScope(func() []fiber.Handler {
		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			func(c *fiber.Ctx) error {
				list, err := con.permission_.AllPermissions(c.UserContext())
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(list),
				)
			},
		}
	})
}

// @Tags        Authorization Management
// @Summary     Create One Route
// @Description Create One Route
// @ID          create-one-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.RouteResponse} "Successfully retrieved Routes"
// @Router      /route [post]
// @Security    Bearer
func (con *AuthorizationController) CreateRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/route").DoWithScope(func() []fiber.Handler {
		body := new(dto.RouteRequest)

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Validate(request.BodyParser(body)),

			func(c *fiber.Ctx) error {
				data, err := con.route_.CreateRoute(c.UserContext(), body)
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
