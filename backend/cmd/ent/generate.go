package main

// _ "github.com/kimchhung/gva/backend/internal/ent/runtime"

//go:generate go run -mod=mod ./entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate go run -mod=mod entgo.io/ent/cmd/ent  generate --feature sql/versioned-migration --target=../../internal/ent ../../app/database/schema --feature intercept --feature sql/execquery --feature sql/modifier --feature schema/snapshot
