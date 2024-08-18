package index

import (
	"context"
	"time"

	"github.com/gva/internal/bootstrap/database"
)

type IndexService struct {
	db *database.Database
}

func NewIndexService(db *database.Database) *IndexService {
	return &IndexService{
		db: db,
	}
}

func (s *IndexService) Now(ctx context.Context) (*time.Time, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT NOW()")
	if err != nil {
		panic(err)
	}

	var now time.Time
	for rows.Next() {
		if err := rows.Scan(&now); err != nil {
			panic(err)
		}
	}

	return &now, nil
}
