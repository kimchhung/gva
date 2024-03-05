package service

import (
	"github.com/kimchhung/gva/extra/app/module/todo/repository"
	"github.com/kimchhung/gva/extra/app/module/todo/dto"

	"github.com/kimchhung/gva/extra/internal/ent"
	"github.com/kimchhung/gva/extra/internal/ent/todo"
	"context"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repository *repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repository,
	}
}

func (s *TodoService) GetTodos(ctx context.Context) ([]*ent.Todo, error) {
	return s.repo.Client().Query().Order(ent.Asc(todo.FieldID)).All(ctx)
}

func (s *TodoService) GetTodoByID(ctx context.Context, id int) (*ent.Todo, error) {
	return s.repo.Client().Query().Where(todo.IDEQ(id)).First(ctx)
}

func (s *TodoService) CreateTodo(ctx context.Context, payload *dto.TodoRequest) (*ent.Todo, error) {
	return s.repo.Client().Create().
		Save(ctx)
}

func (s *TodoService) UpdateTodo(ctx context.Context, id int, payload *dto.TodoRequest) (*ent.Todo, error) {
	return s.repo.Client().UpdateOneID(id).
		Save(ctx)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.Client().DeleteOneID(id).Exec(ctx)
}

