package menu

import (
	"strings"

	"github.com/gva/api/admin/module/menu/dto"
	apperror "github.com/gva/app/common/error"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/utils"
	"github.com/gva/internal/utils/pagi"

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
	return dto.ToMenuResponse(list...), meta, nil
}

func (s *MenuService) GetMenuByID(ctx context.Context, id pxid.ID) (*dto.MenuResponse, error) {
	data, err := s.repo.Q().Where(menu.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToMenuResponse(data)[0], nil
}

func (s *MenuService) CreateMenu(ctx context.Context, r *dto.MenuRequest) (*dto.MenuResponse, error) {
	data, err := s.repo.C().Create().
		SetComponent(r.Component).
		SetPath(r.Path).
		SetNillableIsEnable(r.IsEnable).
		SetMeta(r.Meta).
		SetName(r.Name).
		SetType(r.Type).
		SetNillableParentID(r.Pid).
		SetOrder(r.Order).
		SetNillableRedirect(r.Redirect).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToMenuResponse(data)[0], nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, id pxid.ID, r *dto.MenuRequest) (*dto.MenuResponse, error) {
	update := s.repo.C().UpdateOneID(id).
		SetComponent(r.Component).
		SetPath(r.Path).
		SetNillableIsEnable(r.IsEnable).
		SetMeta(r.Meta).
		SetName(r.Name).
		SetType(r.Type).
		SetOrder(r.Order).
		SetNillableRedirect(r.Redirect)

	if r.Redirect != nil {
		update.SetNillableRedirect(r.Redirect)
	} else {
		update.ClearRedirect()
	}

	if r.Pid != nil {
		isSelfParent := *r.Pid == id
		if isSelfParent {
			return nil, apperror.ErrBadRequest
		}

		update.SetParentID(*r.Pid)
	} else {
		update.ClearPid()
	}

	updated, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToMenuResponse(updated)[0], nil
}

func (s *MenuService) DeleteMenu(ctx context.Context, id pxid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}

func (s *MenuService) EnabledList(ctx context.Context) ([]*dto.MenuResponse, error) {
	list, err := s.repo.Q().Where(menu.IsEnable(true)).All(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToMenuResponse(list...), nil
}
