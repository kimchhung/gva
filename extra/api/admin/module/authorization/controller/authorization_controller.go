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
	_permission *service.PermissionService
	_route      *service.RouteService
}

func (con *AuthorizationController) Init(r fiber.Router) fiber.Router {
	return r.Group("authorization")
}

func NewAuthorizationController(_permission *service.PermissionService, _route *service.RouteService) *AuthorizationController {
	return &AuthorizationController{
		_permission: _permission,
		_route:      _route,
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
				list, meta, err := con._route.Paginate(c.UserContext(), params)
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
				list, err := con._permission.AllPermissions(c.UserContext())
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
// @Summary     Create a Route
// @Description Create a Route
// @ID          create-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=dto.RouteResponse} "Successfully created Routes"
// @Router      /route [post]
// @Security    Bearer
// @Param 		info body dto.RouteRequest true "Route Info"
func (con *AuthorizationController) CreateRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/route").DoWithScope(func() []fiber.Handler {
		body := new(dto.RouteRequest)

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Validate(request.BodyParser(body)),

			func(c *fiber.Ctx) error {
				data, err := con._route.CreateRoute(c.UserContext(), body)
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

// @Tags        Authorization Management
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
func (con *AuthorizationController) UpdateRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Put("/route/:id").DoWithScope(func() []fiber.Handler {
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
				data, err := con._route.UpdateRoute(c.UserContext(), params.ID, body)
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

// @Tags        Authorization Management
// @Summary     Delete a Route
// @Description Delete a Route
// @ID          Delete-a-route
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{} "Successfully Delete Routes"
// @Router      /route [delete]
// @Security    Bearer
// @Param 		id path int true "Route ID"
func (con *AuthorizationController) DeleteRoute(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/route/:id").DoWithScope(func() []fiber.Handler {
		params := new(struct {
			ID int `params:"id" validate:"required,min=0"`
		})

		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c *fiber.Ctx) error {
				err := con._route.DeleteRoute(c.UserContext(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c)
			},
		}
	})
}
