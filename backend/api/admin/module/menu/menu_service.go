package menu

import (
	"strings"

	"github.com/gva/api/admin/module/menu/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"
	"github.com/gva/utils"
	"github.com/gva/utils/pagi"

	"context"

	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/menu"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{
		repo: repo,
	}
}

func (s *MenuService) toDto(value ...*ent.Menu) []*dto.MenuResponse {
	list := make([]*dto.MenuResponse, len(value))
	for i, v := range value {
		list[i] = &dto.MenuResponse{Menu: v}
	}

	return list
}

func (s *MenuService) Paginate(ctx context.Context, p *dto.MenuPagedRequest) ([]*dto.MenuResponse, *pagi.Meta, error) {
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
	listCh := utils.Async(func() []*ent.Menu {
		if !strings.Contains(p.Selects, "list") {
			return nil
		}
		return listQuery.AllX(ctx)
	})
	list := <-listCh
	meta := <-metaCh
	return s.toDto(list...), meta, nil
}

func (s *MenuService) GetMenuByID(ctx context.Context, id xid.ID) (*dto.MenuResponse, error) {
	data, err := s.repo.Q().Where(menu.ID(id)).WithParent().First(ctx)
	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *MenuService) CreateMenu(ctx context.Context, r *dto.MenuRequest) (*dto.MenuResponse, error) {
	data, err := s.repo.C().Create().
		SetComponent(r.Component).
		SetPath(r.Path).
		SetIsEnable(r.IsEnable).
		SetMeta(r.Meta).
		SetName(r.Name).
		SetType(r.Type).
		SetNillableParentID(r.Pid).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, id xid.ID, r *dto.MenuRequest) (*dto.MenuResponse, error) {
	update := s.repo.C().UpdateOneID(id).
		SetComponent(r.Component).
		SetPath(r.Path).
		SetIsEnable(r.IsEnable).
		SetMeta(r.Meta).
		SetName(r.Name).
		SetType(r.Type)

	if r.Pid != nil {
		update.SetParentID(*r.Pid)
	} else {
		update.ClearPid()
	}

	data, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *MenuService) DeleteMenu(ctx context.Context, id xid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
