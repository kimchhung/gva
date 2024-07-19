package todo

import (
	"context"

	"github.com/gva/api/admin/module/todo/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/todo"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repository *repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repository,
	}
}

func (s *TodoService) toDto(value ...*ent.Todo) []*dto.TodoResponse {
	list := make([]*dto.TodoResponse, len(value))
	for i, _ := range value {
		// todo: map value to response value here
		list[i] = &dto.TodoResponse{}
	}
	return list
}

func (s *TodoService) GetTodos(ctx context.Context) ([]*dto.TodoResponse, error) {
	list, err := s.repo.C().Query().Order(ent.Asc(todo.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(list...), nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, id xid.ID) (*dto.TodoResponse, error) {
	data, err := s.repo.C().Query().Where(todo.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(data)[0], nil 
}

func (s *TodoService) CreateTodo(ctx context.Context, payload *dto.TodoRequest) (*dto.TodoResponse, error) {
	create := s.repo.C().Create()
	create.SetName(payload.Name)
	created, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(created)[0], nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, id xid.ID, payload *dto.TodoRequest) (*dto.TodoResponse, error) {
	update := s.repo.C().UpdateOneID(id)
	update.SetName(payload.Name)
	updated, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(updated)[0], nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
