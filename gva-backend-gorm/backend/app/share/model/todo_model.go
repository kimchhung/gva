package model

import "backend/app/share/constant/table"

type Todo struct {
	BaseModel
	Title string `json:"title"`
}

func (t *Todo) TableName() string {
	return table.Todo
}

func (t *Todo) NewTodoModel() *Todo {
	return &Todo{
		BaseModel: NewBaseModel(),
		Title:       t.Title,
	}
}
