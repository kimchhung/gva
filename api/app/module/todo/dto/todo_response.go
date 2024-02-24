package dto

import "github.com/kimchhung/gva/internal/ent"

// Requests & responses Data Transfer Object
type TodoResponse struct {
	*ent.Todo
}
