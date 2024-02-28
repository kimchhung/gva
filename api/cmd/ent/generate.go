package main

import (
	_ "github.com/kimchhung/gva/internal/ent/runtime"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent  generate --feature sql/versioned-migration --target=../../internal/ent ../../app/database/schema --feature intercept
