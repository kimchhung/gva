package operationlog

import (
	"backend/app/admin/middleware"
	adminmiddleware "backend/app/admin/middleware"
	"backend/app/admin/module/operationlog/dto"
	"backend/app/share/permission"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*OperationLogController)(nil)

type OperationLogController struct {
	middleware middleware.Middleware
	service    *OperationLogService
}

func NewOperationLogController(
	middleware *middleware.Middleware,
	service *OperationLogService,
) *OperationLogController {
	return &OperationLogController{
		service:    service,
		middleware: *middleware,
	}
}

func (con *OperationLogController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/operation-log",
			con.middleware.JwtGuard(),
			con.middleware.IpGuard(),
			adminmiddleware.SkipOperationLog(),
		),
	)
}

// @Tags			OperationLog
// @Summary			Get OperationLogs
// @Description		Get OperationLogs
// @ID				get-operation-logs
// @Accept			json
// @Param 			page	query	int	false	"page"
// @Param 			limit	query	int	false	"limit"
// @Param 			search	query	string	false	"search"
// @Param 			filters	query	string	false	"filters"
// @Param 			sorts	query	string	false	"sorts"
// @Param 			orders	query	string	false	"orders"
// @Param 			selects	query	string	false	"selects: list, totalCount"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=[]dto.OperationLogResponse}	"Successfully get OperationLogs"
// @Router			/operation-log [get]
func (con *OperationLogController) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAny(permission.OperationLogView),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.GetOperationLogs(c.Request().Context(), query)
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
