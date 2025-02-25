package blog

import (
	"context"
	"errors"
	"fmt"
	"backend/api/admin/module/blog/dto"
	apperror "backend/app/common/error"
	"backend/app/common/model"
	repository "backend/app/common/repository"
	"backend/internal/bootstrap/database"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/utils"

	"gorm.io/gorm"
)

type BlogService struct {
	repo *repository.BlogRepo
}

// NewBlogService initializes a new BlogService with a JwtService and a UserStore.
func NewBlogService(repo *repository.BlogRepo) *BlogService {
	return &BlogService{
		repo: repo,
	}
}

// CreateBlog creates a new Blog.
func (s *BlogService) CreateBlog(ctx context.Context, p *dto.CreateBlogRequest) (*dto.BlogResponse, error) {
	body := utils.MustCopy(new(model.Blog), p)
	// Default base model
	body.BaseModel = model.NewBaseModel()
	created, err := s.repo.Create(ctx, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.BlogResponse), created), nil
}

// GetBlog gets a Blog by ID.
func (s *BlogService) GetBlog(ctx context.Context, id uint) (*dto.BlogResponse, error) {
	blog, err := s.repo.GetById(ctx, id)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}

		return nil, err
	}

	return utils.MustCopy(new(dto.BlogResponse), blog), nil
}

// LockForUpdate locks a Blog for update.
func (s *BlogService) LockForUpdate(ctx context.Context, id uint) database.TxOperaton {
	return func(tx *gorm.DB) error {
		_, err := s.repo.Tx(tx).GetById(ctx, id, gormq.WithSelect("id"), gormq.WithLockUpdate())
		if err != nil {
			// Check if the error is a not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(apperror.ErrNotFound)
			}

			return err
		}

		return nil
	}
}

// UpdateBlog updates a Blog.
func (s *BlogService) UpdateBlog(ctx context.Context, id uint, p *dto.UpdateBlogRequest) (updatedRes *dto.BlogResponse, err error) {
	err = s.repo.MultiTransaction(
		s.LockForUpdate(ctx, id),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(new(model.Blog), p)
			body.ID = id

			updated, err := s.repo.Tx(tx).Update(ctx, body)
			if err != nil {
				return err
			}

			updatedRes = utils.MustCopy(new(dto.BlogResponse), updated)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return s.GetBlog(ctx, id)
}

// UpdateBlog updates a Blog.
func (s *BlogService) UpdatePatchBlog(ctx context.Context, id uint, p *dto.UpdatePatchBlogRequest) (resp map[string]any, err error) {
	err = s.repo.MultiTransaction(
		s.LockForUpdate(ctx, id),
		func(tx *gorm.DB) error {
			columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
				// allow update status in partial
				"status": gormq.Ignore(),
			})
			dbCols, res := utils.StructToMap(p, columnMap)
			resp = res
			if len(dbCols) == 0 {
				return apperror.NewError(apperror.ErrBadRequest, apperror.Join(
					fmt.Errorf("required at least one field to update, support fields: %s", columnMap.Keys()),
				))
			}

			return tx.Model(&model.Blog{}).
				Scopes(gormq.Equal("id", id)).
				Updates(dbCols).Error
		},
	)
	return
}

// DeleteBlog deletes a Blog by ID.
func (s *BlogService) DeleteBlog(ctx context.Context, id uint) error {
	err := s.repo.DeleteById(ctx, id)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrNotFound
		}
	}

	return nil
}

// GetBlogs gets all Blogs.
func (s *BlogService) GetBlogs(ctx context.Context, query *dto.GetManyQuery) ([]dto.BlogResponse, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.Ignore(),
		"createdAt": gormq.ToSnake(),
	})

	searchColumns := columnMap.Pick(
		"id",
	).Values()

	resp, respMeta := pagi.PrepareResponse[dto.BlogResponse](&query.QueryDto)
	err := s.repo.GetManyAndCount(ctx, &resp, respMeta.TotalCount,
		gormq.WithPageAndLimit(query.Page, query.Limit),
		gormq.WithFilters(query.Filters, columnMap),
		gormq.WithSorts(query.Sorts, columnMap),
		gormq.WithSearch(query.Search, searchColumns...),
	)

	if err != nil {
		return nil, nil, err
	}
	return resp, respMeta, nil
}
