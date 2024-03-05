package repository

import (
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/ent"
)

type TodoRepository struct {
	DB *database.Database
}

func NewTodoRepository(database *database.Database) *TodoRepository {
	return &TodoRepository{
		database,
	}
}

func (r *TodoRepository) Client() *ent.TodoClient {
	return r.DB.Ent.Todo
}

