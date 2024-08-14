package graph

import (
	"entgo.io/contrib/entgql"
	"github.com/gva/api/web/graph/generated"
	"github.com/gva/api/web/graph/resolver"
	"github.com/gva/env"
	"github.com/gva/internal/bootstrap"
	"github.com/gva/internal/bootstrap/database"
	"github.com/rs/zerolog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Server struct {
	echo     *echo.Echo
	db       *database.Database
	resolver *resolver.Resolver

	queryPath      string
	playgroundPath string
}

func NewServer(echo *echo.Echo, cfg *env.Config, db *database.Database, resolver *resolver.Resolver) *Server {
	return &Server{
		echo:     echo,
		db:       db,
		resolver: resolver,

		queryPath:      cfg.API.Web.BasePath + "/query",
		playgroundPath: cfg.API.Web.BasePath + "/playground",
	}
}

func (s *Server) Register(prefix string) {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: s.resolver,
		}),
	)

	// srv.Use(&debug.Tracer{})
	srv.Use(entgql.Transactioner{TxOpener: s.db.Client})
	srv.AddTransport(&transport.Websocket{})

	playground := playground.ApolloSandboxHandler(
		"GraphQL",
		prefix+"/query",
		playground.WithApolloSandboxEndpointIsEditable(true),
	)

	s.echo.POST(prefix+"/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	s.echo.GET(prefix+"/query", func(c echo.Context) error {
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
	fx.Invoke(
		func(gqlserver *Server, cfg *env.Config, b *bootstrap.Bootstrap, log *zerolog.Logger) {
			go func() {
				// wait for bootstrap started
				<-b.Done()
				gqlserver.Register(cfg.API.Web.BasePath)

				host, port := env.ParseAddr(cfg.App.Port)
				if host == "" {
					host = "http://localhost"
				}

				log.Info().
					Str("playground", host+":"+port+gqlserver.playgroundPath).
					Msg("graphql is registered")
			}()
		},
	),
)
