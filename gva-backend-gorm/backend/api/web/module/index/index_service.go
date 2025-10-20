package index

import (
	"context"
	"time"

	"backend/internal/datetime"
)

type IndexService struct {
}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (s *IndexService) Now(ctx context.Context) (time.Time, error) {
	now := datetime.FromContext(ctx)
	return now.ToTime(), nil
}
