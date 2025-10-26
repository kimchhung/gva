package todo

import (
	"context"
	"errors"
	"fmt"
	"backend/app/admin/module/todo/dto"
	coreerror "backend/core/error"
	"backend/app/share/model"
	repository "backend/app/share/repository"
	"backend/internal/gormq"
	"backend/internal/pagi"
	"backend/core/utils"

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
			return nil, coreerror.ErrNotFound
		}

		return nil, err
	}

	return utils.MustCopy(new(dto.TodoResponse), todo), nil
}

// LockForUpdate locks a Todo for update.
func (s *TodoService) LockForUpdate(ctx context.Context, id uint) gormq.Tx {
	return func(tx *gorm.DB) error {
		_, err := s.repo.Tx(tx).GetById(ctx, id, gormq.WithSelect("id"), gormq.WithLockUpdate())
		if err != nil {
			// Check if the error is a not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				panic(coreerror.ErrNotFound)
			}

			return err
		}

		return nil
	}
}

// UpdateTodo updates a Todo.
func (s *TodoService) UpdateTodo(ctx context.Context, id uint, p *dto.UpdateTodoRequest) (updatedRes *dto.TodoResponse, err error) {
	err = s.repo.MultiTransaction(
		s.LockForUpdate(ctx, id),
		func(tx *gorm.DB) error {
			body := utils.MustCopy(new(model.Todo), p)
			body.ID = id

			updated, err := s.repo.Tx(tx).Update(ctx, body)
			if err != nil {
				return err
			}

			updatedRes = utils.MustCopy(new(dto.TodoResponse), updated)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return s.GetTodo(ctx, id)
}

// UpdateTodo updates a Todo.
func (s *TodoService) UpdatePatchTodo(ctx context.Context, id uint, p *dto.UpdatePatchTodoRequest) (resp map[string]any, err error) {
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
				return coreerror.NewError(coreerror.ErrBadRequest, coreerror.AppendMessage(
					fmt.Sprintf("required at least one field to update, support fields: %s", columnMap.Keys()),
				))
			}

			return tx.Model(&model.Todo{}).
				Scopes(gormq.Equal("id", id)).
				Updates(dbCols).Error
		},
	)
	return
}

// DeleteTodo deletes a Todo by ID.
func (s *TodoService) DeleteTodo(ctx context.Context, id uint) error {
	err := s.repo.DeleteById(ctx, id)
	if err != nil {
		// Check if the error is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return coreerror.ErrNotFound
		}
	}

	return nil
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
