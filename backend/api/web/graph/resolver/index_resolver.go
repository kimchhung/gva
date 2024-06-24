package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/kimchhung/gva/backend/api/web/graph/generated"
)

// Now is the resolver for the now field.
func (r *queryResolver) Now(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Now - now"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
