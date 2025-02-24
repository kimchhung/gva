package operationlog

import (
	"backend/api/admin/module/operationlog/dto"
	"backend/app/common/permission"
	"backend/app/common/service"
	"backend/app/middleware"
	"backend/internal/ctr"
	"backend/internal/request"
	"backend/internal/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*OperationLogController)(nil)

type OperationLogController struct {
	service *OperationLogService
	jwt_s   *service.JwtService
}

func NewOperationLogController(service *OperationLogService, jwt_s *service.JwtService) *OperationLogController {
	return &OperationLogController{
		service: service,
		jwt_s:   jwt_s,
	}
}

func (con *OperationLogController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/operation-log", con.jwt_s.RequiredAdmin(), middleware.SkipOperationLogger()),
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
			permission.RequireAnyWithSuper(permission.OperationLogView),
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
