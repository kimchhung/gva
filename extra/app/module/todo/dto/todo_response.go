package dto

import "github.com/kimchhung/gva/extra/internal/ent"

// Requests & responses Data Transfer Object
type TodoResponse struct {
	*ent.Todo
}
