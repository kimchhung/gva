package permission

import (
	"github.com/gofiber/fiber/v2"

	permissions "github.com/kimchhung/gva/extra/app/common/permission"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*PermissionController)(nil)

type PermissionController struct {
	permission_s *PermissionService
}

func (con *PermissionController) Init(r fiber.Router) fiber.Router {
	return r.Group("permission")
}

func NewPermissionController(permission_s *PermissionService) *PermissionController {
	return &PermissionController{
		permission_s: permission_s,
	}
}

// @Tags        Permission
// @Summary     List all permissions
// @Description Get a list of all permissions
// @ID          list-all-permissions
// @Produce     json
// @Success     200 {object} response.Response{data=[]dto.PermissionResponse} "Successfully retrieved Routes"
// @Router      /permission [get]
// @Security    Bearer
func (con *PermissionController) Permissions(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").DoWithScope(func() []fiber.Handler {
		return []fiber.Handler{
			permissions.RequireSuperAdmin(),
			func(c *fiber.Ctx) error {
				list, err := con.permission_s.AllPermissions(c.UserContext())
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
