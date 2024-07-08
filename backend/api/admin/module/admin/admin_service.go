package admin

import (
	"context"

	"github.com/gva/api/admin/module/admin/dto"
	appctx "github.com/gva/app/common/context"
	"github.com/gva/app/common/repository"
	"github.com/gva/app/database/schema/xid"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/role"
	"github.com/gva/internal/ent/route"
	"github.com/gva/utils/pagi"
	"github.com/gva/utils/routeutil"
)

type AdminService struct {
	admin_r *repository.AdminRepository
	db      *database.Database
}

func NewAdminService(
	db *database.Database,
	repo *repository.AdminRepository,
) *AdminService {
	return &AdminService{
		db:      db,
		admin_r: repo,
	}
}

func (s *AdminService) toDto(value ...*ent.Admin) []*dto.AdminResponse {
	list := make([]*dto.AdminResponse, len(value))
	for i, v := range value {
		list[i] = &dto.AdminResponse{
			Admin: v,
		}
	}

	return list
}

func (s *AdminService) Paginate(ctx context.Context, p *dto.AdminPaginateRequest) ([]*dto.AdminResponse, *pagi.Meta, error) {
	query := s.admin_r.Q(
		pagi.WithFilter(p.FilterExp.String(), p.FilterArgs),
		pagi.WithSort(p.Sort...),
		pagi.WithSelect(p.Select...),
	)

	meta := &pagi.Meta{
		Limit:  p.Limit,
		Offset: p.Offset,
	}

	if p.IsCount {
		total := query.CountX(ctx)
		meta.Total = &total
	}

	list := query.Modify(pagi.WithLimitOffset(p.Limit, p.Offset)).AllX(ctx)
	return s.toDto(list...), meta, nil
}

func (s *AdminService) GetAdminByID(ctx context.Context, id xid.ID) (*dto.AdminResponse, error) {
	data, err := s.admin_r.C().Query().Where(admin.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toDto(data)[0], nil
}

func (s *AdminService) CreateAdmin(ctx context.Context, request *dto.AdminRequest) (*dto.AdminResponse, error) {
	data, err := s.admin_r.C().Create().
		SetPassword(request.Password).
		SetUsername(request.Username).
		SetDisplayName(request.DisplayName).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *AdminService) UpdateAdmin(ctx context.Context, id xid.ID, request *dto.AdminRequest) (*dto.AdminResponse, error) {
	data, err := s.admin_r.C().UpdateOneID(id).
		SetDisplayName(request.DisplayName).
		SetUsername(request.Username).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return s.toDto(data)[0], nil
}

func (s *AdminService) DeleteAdmin(ctx context.Context, id xid.ID) error {
	return s.admin_r.C().DeleteOneID(id).Exec(ctx)
}

func (s *AdminService) GetAdminNestedRouteById(ctx context.Context, adminId xid.ID) ([]*ent.Menu, error) {
	if appctx.MustAdminContext(ctx).IsSuperAdmin() {
		routes, err := s.db.Menu.Query().Where(route.IsEnable(true)).All(ctx)
		if err != nil {
			return nil, err
		}

		return routeutil.GroupRouteToNested(routes), nil
	}

	routes, err := s.db.Role.Query().
		Where(role.HasAdminsWith(admin.ID(adminId))).
		QueryRoutes().
		Where(route.IsEnable(true)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return routeutil.GroupRouteToNested(routes), nil
}

func (s *AdminService) GetAdminPermissionById(ctx context.Context, adminId xid.ID) ([]*ent.Permission, error) {
	routes, err := s.db.Role.Query().
		Where(
			role.HasAdminsWith(admin.ID(adminId)),
		).
		QueryPermissions().All(ctx)

	if err != nil {
		return nil, err
	}

	return routes, nil
}
