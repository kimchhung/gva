package module_template

var Service = `package {{.EntityAllLower}}

import (
	"context"
	"errors"
	"fmt"
	"backend/api/admin/module/{{.EntityAllLower}}/dto"
	apperror "backend/app/common/error"
	"backend/app/common/model"
	repository "backend/app/common/repository"
	"backend/internal/bootstrap/database"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/utils"

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
	body.BaseModelV2 = model.NewBaseModelV2()
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

// LockForUpdate locks a {{.EntityPascal}} for update.
func (s *{{.EntityPascal}}Service) LockForUpdate(ctx context.Context, id uint) database.TxOperaton {
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

// Update{{.EntityPascal}} updates a {{.EntityPascal}}.
func (s *{{.EntityPascal}}Service) Update{{.EntityPascal}}(ctx context.Context, id uint, p *dto.Update{{.EntityPascal}}Request) (updatedRes *dto.{{.EntityPascal}}Response, err error) {
	err = s.repo.MultiTransaction(
		s.LockForUpdate(ctx, id),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(new(model.{{.EntityPascal}}), p)
			body.ID = id

			updated, err := s.repo.Tx(tx).Update(ctx, body)
			if err != nil {
				return err
			}

			updatedRes = utils.MustCopy(new(dto.{{.EntityPascal}}Response), updated)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return s.Get{{.EntityPascal}}(ctx, id)
}

// Update{{.EntityPascal}} updates a {{.EntityPascal}}.
func (s *{{.EntityPascal}}Service) UpdatePatch{{.EntityPascal}}(ctx context.Context, id uint, p *dto.UpdatePatch{{.EntityPascal}}Request) (resp map[string]any, err error) {
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

			return tx.Model(&model.GameSource{}).
				Scopes(gormq.Equal("id", id)).
				Updates(dbCols).Error
		},
	)
	return
}

// Delete{{.EntityPascal}} deletes a {{.EntityPascal}} by ID.
func (s *{{.EntityPascal}}Service) Delete{{.EntityPascal}}(ctx context.Context, id uint) error {
	err := s.repo.DeleteById(ctx, id)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrNotFound
		}
	}

	return nil
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
