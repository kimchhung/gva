package dto

import (
	"backend/app/share/constant/table"
	"backend/app/share/model"
)

// Requests & responses Data Transfer Object
type TodoResponse struct {
	model.BaseModel

	Title string `json:"title"`
}

func (TodoResponse) TableName() string {
	return table.Todo
}
