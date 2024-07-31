package department

import (
	"context"
	"strings"

	"github.com/gva/api/admin/module/department/dto"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/department"
	"github.com/gva/utils"
	"github.com/gva/utils/pagi"
)

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repo: repo,
	}
}

func (s *DepartmentService) toDto(value ...*ent.Department) []*dto.DepartmentResponse {
	list := make([]*dto.DepartmentResponse, len(value))
	for i, v := range value {
		// todo: map value to response value here
		list[i] = &dto.DepartmentResponse{
			Department: v,
		}
	}
	return list
}

func (s *DepartmentService) GetDepartments(ctx context.Context, p *dto.DepartmentPagedRequest) ([]*dto.DepartmentResponse, *pagi.Meta, error) {
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
	listCh := utils.Async(func() []*ent.Department {
		if !strings.Contains(p.Selects, "list") {
			return nil
		}
		return listQuery.AllX(ctx)
	})
	list := <-listCh
	meta := <-metaCh
	return s.toDto(list...), meta, nil
}

func (s *DepartmentService) GetDepartmentByID(ctx context.Context, id pxid.ID) (*dto.DepartmentResponse, error) {
	data, err := s.repo.C().Query().Where(department.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(data)[0], nil
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, payload *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	create := s.repo.C().Create()
	create.SetName(payload.Name).
		SetName(payload.Name).
		SetNameID(payload.NameId).
		SetNillableIsEnable(payload.IsEnable).
		SetPid(*payload.Pid)

	created, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(created)[0], nil
}

func (s *DepartmentService) UpdateDepartment(ctx context.Context, id pxid.ID, payload *dto.DepartmentRequest) (*dto.DepartmentResponse, error) {
	update := s.repo.C().UpdateOneID(id)
	update.SetName(payload.Name).
		SetName(payload.Name).
		SetNameID(payload.NameId).
		SetNillableIsEnable(payload.IsEnable).
		SetPid(*payload.Pid)

	updated, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(updated)[0], nil
}

func (s *DepartmentService) DeleteDepartment(ctx context.Context, id pxid.ID) error {
	return s.repo.C().DeleteOneID(id).Exec(ctx)
}
