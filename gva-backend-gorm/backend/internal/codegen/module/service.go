package module_template

var Service = `package {{.EntityAllLower}}

import (
	"context"
	"errors"
	"fmt"
	"backend/app/admin/module/{{.EntityAllLower}}/dto"
	apperror "backend/app/share/error"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/core/utils"

	"gorm.io/gorm"
)

type {{.EntityPascal}}Service struct {
	repo *repository.{{.EntityPascal}}Repo
}

// New{{.EntityPascal}}Service initializes a new {{.EntityPascal}}Service with a JwtService and a UserStore.
func New{{.EntityPascal}}Service(repo *repository.{{.EntityPascal}}Repo) *{{.EntityPascal}}Service {
	return &{{.EntityPascal}}Service{
		repo: repo,
	}
}

// Create{{.EntityPascal}} creates a new {{.EntityPascal}}.
func (s *{{.EntityPascal}}Service) Create{{.EntityPascal}}(ctx context.Context, p *dto.Create{{.EntityPascal}}Request) (*dto.{{.EntityPascal}}Response, error) {
	body := utils.MustCopy(new(model.{{.EntityPascal}}), p)
	// Default base model
	body.BaseModel = model.NewBaseModel()
	created, err := s.repo.Create(ctx, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.{{.EntityPascal}}Response), created), nil
}

// Get{{.EntityPascal}} gets a {{.EntityPascal}} by ID.
func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}(ctx context.Context, id uint) (*dto.{{.EntityPascal}}Response, error) {
	{{.EntityAllLower}}, err := s.repo.GetById(ctx, id)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}

		return nil, err
	}

	return utils.MustCopy(new(dto.{{.EntityPascal}}Response), {{.EntityAllLower}}), nil
}

// lockForUpdate locks a {{.EntityPascal}} for update.
func (s *{{.EntityPascal}}Service) lockForUpdate(ctx context.Context, id uint, out *model.{{.EntityPascal}}, opts ...gormq.Option) gormq.Tx {
	return func(tx *gorm.DB) error {
		target, err := s.repo.Tx(tx).GetById(ctx, id, gormq.Multi(opts...), gormq.WithLockUpdate())
		if err != nil {
			// Check if the error is a not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.ErrNotFound
			}

			return err
		}

		if out != nil && target != nil {
			*out = *target
		}

		return nil
	}
}

// Update{{.EntityPascal}} updates a {{.EntityPascal}}.
func (s *{{.EntityPascal}}Service) Update{{.EntityPascal}}(ctx context.Context, id uint, dtoReq *dto.Update{{.EntityPascal}}Request) (updatedRes *dto.{{.EntityPascal}}Response, err error) {
	var (
		target model.{{.EntityPascal}}
	)

	err = s.repo.MultiTransaction(
		s.lockForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(&target, dtoReq)
			updated, err := s.repo.Tx(tx).Update(ctx, body)
			if err != nil {
				return err
			}

			updatedRes = utils.MustCopy(new(dto.{{.EntityPascal}}Response), updated)
			return nil
		},
	)

	return
}

// Update{{.EntityPascal}} updates a {{.EntityPascal}}.
func (s *{{.EntityPascal}}Service) UpdatePatch{{.EntityPascal}}(ctx context.Context, id uint, dtoReq *dto.UpdatePatch{{.EntityPascal}}Request) (resp map[string]any, err error) {
	var (
		target model.{{.EntityPascal}}
	)

	err = s.repo.MultiTransaction(
		s.lockForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
				// allow update status in partial
				"status": gormq.Ignore(),
			})

			dbCols, res := utils.StructToMap(dtoReq, columnMap)
			if len(dbCols) == 0 {
				return apperror.NewError(apperror.ErrBadRequest, apperror.AppendMessage(
					fmt.Sprintf("required at least one field to update, support fields: %s", columnMap.Keys()),
				))
			}

			if err := tx.Model(&target).
				Scopes(gormq.Equal("id", id)).
				Updates(dbCols).Error; err != nil {
				return err
			}

			resp = res
			return nil
		},
	)

	return
}

// Delete{{.EntityPascal}} deletes a {{.EntityPascal}} by ID.
func (s *{{.EntityPascal}}Service) Delete{{.EntityPascal}}(ctx context.Context, id uint) error {
	var (
		target model.{{.EntityPascal}}
	)

	return s.repo.MultiTransaction(s.lockForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			return s.repo.DeleteById(ctx, id)
		},
	)
}

// Get{{.EntityPascal}}s gets all {{.EntityPascal}}s.
func (s *{{.EntityPascal}}Service) Get{{.EntityPascal}}s(ctx context.Context, query *dto.GetManyQuery) ([]dto.{{.EntityPascal}}Response, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.Ignore(),
		"createdAt": gormq.ToSnake(),
	})

	searchColumns := columnMap.Pick(
		"id",
	).Values()

	resp, respMeta := pagi.PrepareResponse[dto.{{.EntityPascal}}Response](&query.QueryDto)
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
`
