package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type CreateTodoRequest struct {
	Title string `json:"title" validate:"required"`
}

type GetTodoRequest struct {
	ID uint `param:"id" validate:"required"`
}

type UpdateTodoRequest struct {
	Title string `json:"title"`
}


// not nil will update
type UpdatePatchTodoRequest struct {
	Title *string `json:"title"`
}

type GetManyQuery struct {
	pagi.QueryDto
}
