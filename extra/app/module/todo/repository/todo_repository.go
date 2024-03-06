package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
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

func (r *TodoRepository) DB() *database.Database {
	return r.db
}

