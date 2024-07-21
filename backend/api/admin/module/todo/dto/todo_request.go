package dto
import "github.com/gva/internal/rql"

// Requests Data Transfer Object
type TodoRequest struct {
	Name string `json:"name"`
}

type TodoPagedRequest struct {
	rql.Params
	Selects string `query:"selects"`
}
