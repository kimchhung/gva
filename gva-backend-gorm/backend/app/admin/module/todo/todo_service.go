package todo

import (
	"backend/app/admin/module/todo/dto"
	apperror "backend/app/share/error"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	"backend/core/utils"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TodoService struct {
	repo *repository.TodoRepo
}

// NewTodoService initializes a new TodoService with a JwtService and a UserStore.
func NewTodoService(repo *repository.TodoRepo) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

// CreateTodo creates a new Todo.
func (s *TodoService) CreateTodo(ctx context.Context, p *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
	body := utils.MustCopy(new(model.Todo), p)
	// Default base model
	body.BaseModel = model.NewBaseModel()
	created, err := s.repo.Create(ctx, body)
	if err != nil {
		return nil, err
	}

	return utils.MustCopy(new(dto.TodoResponse), created), nil
}

// GetTodo gets a Todo by ID.
func (s *TodoService) GetTodo(ctx context.Context, id uint) (*dto.TodoResponse, error) {
	todo, err := s.repo.GetById(ctx, id)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}

		return nil, err
	}

	return utils.MustCopy(new(dto.TodoResponse), todo), nil
}

// lockForUpdate locks a Todo for update.
func (s *TodoService) lockForUpdate(ctx context.Context, id uint, out *model.Todo, opts ...gormq.Option) gormq.Tx {
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

// UpdateTodo updates a Todo.
func (s *TodoService) UpdateTodo(ctx context.Context, id uint, dtoReq *dto.UpdateTodoRequest) (updatedRes *dto.TodoResponse, err error) {
	var (
		target model.Todo
	)

	err = s.repo.MultiTransaction(
		s.lockForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(&target, dtoReq)
			updated, err := s.repo.Tx(tx).Update(ctx, body)
			if err != nil {
				return err
			}

			updatedRes = utils.MustCopy(new(dto.TodoResponse), updated)
			return nil
		},
	)

	return
}

// UpdateTodo updates a Todo.
func (s *TodoService) UpdatePatchTodo(ctx context.Context, id uint, dtoReq *dto.UpdatePatchTodoRequest) (resp map[string]any, err error) {
	var (
		target model.Todo
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
				return apperror.ErrBadRequest.Copy(
					apperror.DisableTranslate(),
					apperror.AppendMessage(fmt.Sprintf("required at least one field to update, support fields: %s", columnMap.Keys())),
				)
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

// DeleteTodo deletes a Todo by ID.
func (s *TodoService) DeleteTodo(ctx context.Context, id uint) error {
	var (
		target model.Todo
	)

	return s.repo.MultiTransaction(s.lockForUpdate(ctx, id, &target),
		func(tx *gorm.DB) error {
			return s.repo.DeleteById(ctx, id)
		},
	)
}

// GetTodos gets all Todos.
func (s *TodoService) GetTodos(ctx context.Context, query *dto.GetManyQuery) ([]dto.TodoResponse, *pagi.MetaDto, error) {
	columnMap := gormq.MapTableColumn(map[string]gormq.MapOption{
		"id":        gormq.Ignore(),
		"createdAt": gormq.ToSnake(),
	})

	searchColumns := columnMap.Pick(
		"id",
	).Values()

	resp, respMeta := pagi.PrepareResponse[dto.TodoResponse](&query.QueryDto)
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
