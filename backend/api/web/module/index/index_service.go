package index

import (
	"context"
	"time"

	"github.com/kimchhung/gva/backend-echo/internal/bootstrap/database"
)

type IndexService struct {
	db *database.Database
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewIndexService(db *database.Database) *IndexService {
	return &IndexService{
		db: db,
	}
}

func (s *IndexService) Now(ctx context.Context) (time.Time, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT NOW()")
	if err != nil {
		panic(err)
	}

	var now time.Time

	if rows.Next() {
		if err := rows.Scan(&now); err != nil {
			panic(err)
		}
	}

	return now, nil
}
