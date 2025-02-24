package dto

import (
	"backend/internal/pagi"
)

// Requests & responses Data Transfer Object

type GetOperationLogRequest struct {
	ID uint `param:"id" validate:"required"`
}

type GetManyQuery struct {
	pagi.QueryDto
}
