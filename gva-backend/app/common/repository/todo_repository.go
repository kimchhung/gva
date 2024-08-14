package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/utils/pagi"
)

type TodoRepository struct {
	db *database.Database
}

func NewTodoRepository(database *database.Database) *TodoRepository {
	return &TodoRepository{
		database,
	}
}

func (r *TodoRepository) C() *ent.TodoClient {
	return r.db.Todo
}

// For query
func (r *TodoRepository) Q(opts ...pagi.InterceptorOption) *ent.TodoQuery {
	if len(opts) == 0 {
		return r.C().Query()
	}

	return pagi.WithInterceptor(r.C().Query(), opts...)
}

