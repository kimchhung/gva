package middleware

import (
	"backend/app/share/repository"
	"backend/app/share/service"
	"backend/core/database"
	coretype "backend/core/type"

	"go.uber.org/zap"
)

var _ interface {
	coretype.Middleware
} = (*Middleware)(nil)

type Middleware struct {
	db      *database.Database
	log     *zap.Logger
	ip_s    *service.IPService
	jwt_s   *service.JwtService
	admin_r *repository.AdminRepo
}

func NewMiddleware(
	db *database.Database,
	log *zap.Logger,
	ip_s *service.IPService,
	jwt_s *service.JwtService,
	admin_r *repository.AdminRepo,
) *Middleware {
	return &Middleware{
		db:      db,
		log:     log,
		ip_s:    ip_s,
		jwt_s:   jwt_s,
		admin_r: admin_r,
	}
}

func (m *Middleware) RegisterMiddleware(c coretype.MiddlewareRouter) {
	c.Use(
		m.OperationLog(),
	)
}
