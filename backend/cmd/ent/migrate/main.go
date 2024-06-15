//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kimchhung/gva/backend/internal/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	name, env := os.Args[1], os.Args[2]
	mapConfig, err := godotenv.Read(fmt.Sprintf("app/database/%v.connection.env", env))
	if err != nil {
		log.Fatalf("failed to load connection config %v", err)
	}

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
	err = migrate.NamedDiff(context.Background(), mapConfig["devUrl"], name, opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}

	// uncomment if want to generate sql for data migration
	// migratedata.GenerateSql(dir)
}
