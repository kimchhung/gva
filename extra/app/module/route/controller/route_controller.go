package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/app/common/permissions"
	"github.com/kimchhung/gva/extra/app/common/services"
	"github.com/kimchhung/gva/extra/app/module/route/dto"
	"github.com/kimchhung/gva/extra/app/module/route/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*RouteController)(nil)

type RouteController struct {
	service    *service.RouteService
	jwtService *services.JwtService
}

func (con *RouteController) Routes(r fiber.Router) {
	route := r.Group("route")
	route.Use(con.jwtService.ProtectAdmin())
	rctrl.Register(route, con)
}

func NewRouteController(service *service.RouteService, jwtService *services.JwtService) *RouteController {
	return &RouteController{
		service:    service,
		jwtService: jwtService,
	}
}

// @Tags Route
// @Summary List all Routes
// @Description Get a list of all Routes
// @ID list-all-Routes
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.RouteResponse{list=[]dto.RouteResponse}} "Successfully retrieved Routes"
// @Router /route [get]
// @Security Bearer
func (con *RouteController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Routes").Do(
		permissions.RequireSuperAdmin(),
		func(c *fiber.Ctx) error {
			list, err := con.service.GetNestedRoutes(c.UserContext())
			if err != nil {
				return err
			}

			return request.Response(c,
				response.Data(list),
				response.Message("Route list retreived successfully!"),
			)
		},
	)
}

// @Tags Route
// @Security Bearer
// @Summary Get a Route
// @Description Get a Route by ID
// @ID get-Route-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Route ID"
// @Success   200 {object} response.Response{data=dto.RouteResponse}
// @Router /route/{id} [get]
func (con *RouteController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Route").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.GetRouteByID(c.UserContext(), param.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The route retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Route
// @Security Bearer
// @Summary Create a Route
// @Description Create a new Route with the provided details
// @ID create-Route
// @Accept  json
// @Produce  json
// @Param Route body dto.RouteRequest true "Route data"
// @Success  200 {object} response.Response{data=dto.RouteResponse} "Successfully created Route"
// @Router /route [post]
func (con *RouteController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Route").DoWithScope(func() []fiber.Handler {
		body := new(dto.RouteRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreateRoute(c.UserContext(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The route retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Route
// @Security Bearer
// @Summary Update a Route
// @Description Update a Route by ID
// @ID update-Route-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Route ID"
// @Param Route body dto.RouteRequest true "Route data"
// @Success  200 {object} response.Response{data=dto.RouteResponse} "Successfully updated Route"
// @Router /route/{id} [patch]
func (con *RouteController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Route").DoWithScope(func() []fiber.Handler {
		body := new(dto.RouteRequest)
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.UpdateRoute(c.UserContext(), param.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The route retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Route
// @Security Bearer
// @Summary Delete a Route
// @Description Delete a Route by ID
// @ID delete-Route-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Route ID"
// @Success  200 {object} response.Response{} "Successfully deleted Route"
// @Router /route/{id} [delete]
func (con *RouteController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Route").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.DeleteRoute(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The route retrieved successfully!"),
				)
			},
		}
	})
}
