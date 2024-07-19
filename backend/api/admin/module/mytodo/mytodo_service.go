package mytodo

import (
	"github.com/gva/app/common/repository"
	"github.com/gva/api/admin/module/mytodo/dto"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/mytodo"
	"context"
)

type MyTodoService struct {
	repo *repository.MyTodoRepository
}

func NewMyTodoService(repository *repository.MyTodoRepository) *MyTodoService {
	return &MyTodoService{
		repo: repository,
	}
}

func (s *MyTodoService) GetMyTodos(ctx context.Context) ([]*ent.MyTodo, error) {
	return s.repo.C().Query().Order(ent.Asc(mytodo.FieldID)).All(ctx)
}

func (s *MyTodoService) GetMyTodoByID(ctx context.Context, id xid.ID) (*ent.MyTodo, error) {
	return s.repo.C().Query().Where(mytodo.IDEQ(id)).First(ctx)
}

func (s *MyTodoService) CreateMyTodo(ctx context.Context, payload *dto.MyTodoRequest) (*ent.MyTodo, error) {
	return s.repo.C().Create().
		Save(ctx)
}

func (s *MyTodoService) UpdateMyTodo(ctx context.Context, id xid.ID, payload *dto.MyTodoRequest) (*ent.MyTodo, error) {
	return s.repo.C().UpdateOneID(id).
		Save(ctx)
}

func (s *MyTodoService) DeleteMyTodo(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}

