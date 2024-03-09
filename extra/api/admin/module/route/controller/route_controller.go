package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/api/admin/module/route/service"
	"github.com/kimchhung/gva/extra/app/common/permissions"
	"github.com/kimchhung/gva/extra/app/common/services"
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
