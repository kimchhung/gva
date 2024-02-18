package module

import "github.com/kimchhung/gva/internal/rctrl"

var _ interface {
	rctrl.Router
} = (*Router)(nil)

type Router struct {
	moduleRouters []rctrl.Router
}

func NewRouter(moduleRouters []rctrl.Router) *Router {
	return &Router{moduleRouters}
}

func (r *Router) Register() {
	for _, mr := range r.moduleRouters {
		mr.Register()
	}
}
