package resolver

import (
	"github.com/gva/internal/bootstrap/database"
)

type Resolver struct {
	db *database.Database
}

func NewResolver(db *database.Database) *Resolver {
	return &Resolver{db: db}
}
