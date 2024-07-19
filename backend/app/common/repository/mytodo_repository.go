package repository

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
)

type MyTodoRepository struct {
	db *database.Database
}

func NewMyTodoRepository(database *database.Database) *MyTodoRepository {
	return &MyTodoRepository{
		database,
	}
}

func (r *MyTodoRepository) C() *ent.MyTodoClient {
	return r.db.MyTodo
}

func (r *MyTodoRepository) DB() *database.Database {
	return r.db
}

