package dto

import (
	"github.com/gva/internal/ent"
)

// Responses Data Transfer Object
type DepartmentResponse struct {
	*ent.Department
}
