package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
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

