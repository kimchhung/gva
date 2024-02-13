package service

import (
	"context"
	"gva/app/module/article/dto"
	"gva/app/module/article/repository"
	"gva/internal/ent"
	"gva/internal/ent/article"
)

type ArticleService struct {
	repo *repository.ArticleRepository
}

func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) GetArticles(ctx context.Context) ([]*ent.Article, error) {
	return s.repo.Client().Query().Order(ent.Asc(article.FieldID)).All(ctx)
}

func (s *ArticleService) GetArticleByID(ctx context.Context, id int) (*ent.Article, error) {
	return s.repo.Client().Query().Where(article.IDEQ(id)).First(ctx)
}

func (s *ArticleService) CreateArticle(ctx context.Context, request dto.ArticleRequest) (*ent.Article, error) {
	return s.repo.Client().Create().
		SetTitle(request.Title).
		SetContent(request.Content).
		Save(ctx)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, id int, request dto.ArticleRequest) (*ent.Article, error) {
	return s.repo.Client().UpdateOneID(id).
		SetTitle(request.Title).
		SetContent(request.Content).
		Save(ctx)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}
