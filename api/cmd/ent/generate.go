package main

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --target=../../internal/ent ../../app/database/schema
