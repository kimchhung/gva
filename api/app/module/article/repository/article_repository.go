package repository

import (
	"gva/internal/bootstrap/database"
	"gva/internal/ent"
)

type ArticleRepository struct {
	DB *database.Database
}

func NewArticleRepository(database *database.Database) *ArticleRepository {
	return &ArticleRepository{
		DB: database,
	}
}

func (repo *ArticleRepository) Client() *ent.ArticleClient {
	return repo.DB.Ent.Article
}
