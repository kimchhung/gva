//go:build ignore

package main

import (
	"context"
	"fmt"
	"gva/cmd/ent/migrate/migratedata"
	"gva/internal/ent/migrate"
	"log"
	"os"
	"strings"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type SeedFunc = func(dir *atlas.LocalDir) error

func runSeeds(dir *atlas.LocalDir, seedFuncs ...SeedFunc) {
	for _, sf := range seedFuncs {
		if err := sf(dir); err != nil {
			log.Fatalf("failed to generate seed")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	mapConfig, err := godotenv.Read(fmt.Sprintf("app/database/%v.connection.env", os.Args[2]))
	if err != nil {
		log.Fatalf("failed to load connection config %v", err)
	}

	fmt.Printf("config :%v", os.Args[1])
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir(strings.Replace(mapConfig["dir"], "file://", "", 1))
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(os.Args) < 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, mapConfig["devUrl"], os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}

	runSeeds(dir, migratedata.SeedAdmin)
}
