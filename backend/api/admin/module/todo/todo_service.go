package todo

import (
	"context"
	"strings"

	"github.com/gva/api/admin/module/todo/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/todo"
	"github.com/gva/utils"
	"github.com/gva/utils/pagi"
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
	for i, v := range value {
		// todo: map value to response value here
		list[i] = &dto.TodoResponse{v}
	}
	return list
}

func (s *TodoService) GetTodos(ctx context.Context, p *dto.TodoPagedRequest) ([]*dto.TodoResponse, *pagi.Meta, error) {
	if p.Selects == "" {
		p.Selects = "count,list"
	}
	query := s.repo.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	)
	countQuery := query.Clone()
	listQuery := query.Modify(pagi.WithLimitOffset(p.Limit, p.Offset))
	metaCh := utils.Async(func() *pagi.Meta {
		m := &pagi.Meta{Limit: p.Limit, Offset: p.Offset}
		if !strings.Contains(p.Selects, "count") {
			return m
		}
		m.Total = countQuery.CountX(ctx)
		return m
	})
	listCh := utils.Async(func() []*ent.Todo {
		if !strings.Contains(p.Selects, "list") {
			return nil
		}
		return listQuery.AllX(ctx)
	})
	list := <-listCh
	meta := <-metaCh
	return s.toDto(list...), meta, nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, id xid.ID) (*dto.TodoResponse, error) {
	data, err := s.repo.Q().Where(todo.IDEQ(id)).First(ctx)
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
