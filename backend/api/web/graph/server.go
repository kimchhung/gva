package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kimchhung/gva/backend/api/web/graph/resolver"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Register() {
	srv := handler.NewDefaultServer(resolver.NewSchema())
	{
		s.echo.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		s.echo.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

}
