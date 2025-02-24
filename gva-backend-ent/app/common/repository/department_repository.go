package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/utils/pagi"
)

type DepartmentRepository struct {
	db *database.Database
}

func NewDepartmentRepository(database *database.Database) *DepartmentRepository {
	return &DepartmentRepository{
		database,
	}
}

func (r *DepartmentRepository) C() *ent.DepartmentClient {
	return r.db.Department
}

// For query
func (r *DepartmentRepository) Q(opts ...pagi.InterceptorOption) *ent.DepartmentQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}
