package graph

import (
	"entgo.io/contrib/entgql"
	"github.com/gva/api/web/graph/generated"
	"github.com/gva/api/web/graph/resolver"
	"github.com/gva/internal/bootstrap/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Server struct {
	echo     *echo.Echo
	db       *database.Database
	resolver *resolver.Resolver
}

func NewServer(echo *echo.Echo, db *database.Database, resolver *resolver.Resolver) *Server {
	return &Server{
		echo:     echo,
		db:       db,
		resolver: resolver,
	}
}

func (s *Server) Register(prefix string) {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: s.resolver,
		}),
	)

	srv.Use(&debug.Tracer{})
	srv.Use(entgql.Transactioner{TxOpener: s.db.Client})

	playground := playground.ApolloSandboxHandler(
		"GraphQL",
		prefix+"/query",
	)

	s.echo.POST(prefix+"/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	s.echo.GET(prefix+"/playground", func(c echo.Context) error {
		playground.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}

var Module = fx.Module("graph-module",
	fx.Provide(resolver.NewResolver),
	fx.Provide(NewServer),
)
