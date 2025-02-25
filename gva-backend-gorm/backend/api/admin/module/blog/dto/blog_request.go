package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object
type CreateBlogRequest struct {
	Title string `json:"title" validate:"required"`
}

type GetBlogRequest struct {
	ID uint `param:"id" validate:"required"`
}

type UpdateBlogRequest struct {
	Title string `json:"title"`
}


// not nil will update
type UpdatePatchBlogRequest struct {
	Title *string `json:"title"`
}

type GetManyQuery struct {
	pagi.QueryDto
}
