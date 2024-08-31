//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/gva/internal/ent/runtime"
)
