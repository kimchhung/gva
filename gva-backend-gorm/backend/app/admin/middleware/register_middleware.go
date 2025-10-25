package adminmiddleware

import (
	coretype "backend/core/type"
)

var _ interface {
	coretype.Middleware
} = (*Middleware)(nil)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) RegisterMiddleware(c coretype.MiddlewareRouter) {
	c.Use(
		m.OperationLog(),
	)
}
