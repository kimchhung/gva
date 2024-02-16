package module

import (
	"github.com/kimchhung/gva/internal/control_route"
)

var _ interface {
	control_route.Router
} = (*Router)(nil)

type Router struct {
	moduleRouters []control_route.Router
}

func NewRouter(moduleRouters []control_route.Router) *Router {
	return &Router{moduleRouters}
}

func (r *Router) Register() {
	for _, mr := range r.moduleRouters {
		mr.Register()
	}
}
