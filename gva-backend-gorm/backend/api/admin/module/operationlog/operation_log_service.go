package operationlog

import (
	"context"

	"backend/api/admin/module/operationlog/dto"
	repository "backend/app/common/repository"
	"backend/internal/gormq"
	"backend/internal/pagi"

	"gorm.io/gorm"
)

type OperationLogService struct {
	repo *repository.OperationLogRepo
}

// NewOperationLogService initializes a new OperationLogService with a JwtService and a UserStore.
func NewOperationLogService(repo *repository.OperationLogRepo) *OperationLogService {
	return &OperationLogService{
		repo: repo,
	}
}

// GetOperationLogs gets all OperationLogs.
func (s *OperationLogService) GetOperationLogs(ctx context.Context, query *dto.GetManyQuery) ([]dto.OperationLogResponse, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"operator":      gormq.ToColumn("Admin.username"),
		"operatingTime": gormq.ToColumn("operation_logs.created_at"),
		"latency":       gormq.ToColumn("operation_logs.latency"),
		"action":        gormq.ToColumn("operation_logs.scope"),
		"code":          gormq.ToColumn("operation_logs.code"),
		"ip":            gormq.ToColumn("operation_logs.ip"),
		"method":        gormq.ToColumn("operation_logs.method"),
		"path":          gormq.ToColumn("operation_logs.path"),
	})

	searchColumns := columnMap.Pick(
		"operator",
		"path",
	).Values()

	query.DefaultSort("-createdAt")

	resp, respMeta := pagi.PrepareResponse[dto.OperationLogResponse](&query.QueryDto)
	err := s.repo.GetManyAndCount(ctx, &resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.WithFilters(query.Filters, columnMap),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
		func(q *gorm.DB) *gorm.DB {
			return q.Joins("Admin")
		},
	)

	respMeta.UpdatePagination(len(resp))

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}
