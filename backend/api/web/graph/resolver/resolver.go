package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/kimchhung/gva/backend/api/web/graph/generated"
)

type Resolver struct{}

func NewSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{},
	})
}
