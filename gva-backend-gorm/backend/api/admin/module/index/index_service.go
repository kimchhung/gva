package index

import (
	"context"
	"time"

	"backend/internal/bootstrap/database"
	"backend/internal/datetime"
)

type IndexService struct {
	db *database.Database
}

func NewIndexService(db *database.Database) *IndexService {
	return &IndexService{
		db: db,
	}
}

func (s *IndexService) Now(ctx context.Context) (time.Time, error) {
	now := datetime.Must()
	return now.ToTime(), nil
}
