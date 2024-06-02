package permission

import (
	"github.com/labstack/echo/v4"

	permissions "github.com/kimchhung/gva/extra/app/common/permission"
	"github.com/kimchhung/gva/extra/internal/echoc"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

// don't remove for runtime type checking
var _ interface{ echoc.Controller } = (*PermissionController)(nil)

type PermissionController struct {
	permission_s *PermissionService
}

func (con *PermissionController) Init(r *echo.Group) *echo.Group {
	return r.Group("/permissions")
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
// @Router      /permissions [get]
// @Security    Bearer
func (con *PermissionController) Permissions(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/").DoWithScope(func() []echo.HandlerFunc {
		return []echo.HandlerFunc{
			permissions.OnlySuperAdmin(),
			func(c echo.Context) error {
				list, err := con.permission_s.AllPermissions(c.Request().Context())
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
