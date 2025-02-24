package comic

import (
	"github.com/gva/internal/bootstrap/database"
)

type ComicService struct {
	db *database.Database
}

func NewComicService(db *database.Database) *ComicService {
	return &ComicService{
		db: db,
	}
}

func (s *ComicService) Fetch() {

}
