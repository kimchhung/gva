//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(true),
		entgql.WithConfigPath("../../gqlgen.yml"),
		entgql.WithSchemaPath("../../api/web/graph/schema/ent.gql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate(
		"../../app/database/schema",
		&gen.Config{
			Target:   "internal/ent",
			Schema:   "app/database/schema",
			Features: gen.AllFeatures,
			Package:  "github.com/gva/internal/ent",
		},
		entc.Extensions(ex),
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	//go:generate go run -mod=mod entgo.io/ent/cmd/ent
	// generate --target=../../internal/ent ../../app/database/schema
	// --feature sql/versioned-migration,intercept,sql/execquery,sql/modifier,schema/snapshot,privacy,entql

}
