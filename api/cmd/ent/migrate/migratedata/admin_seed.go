package migratedata

import (
	"context"
	"fmt"
	"gva/internal/ent"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

// SeedUsers add the initial users to the database.
func SeedSuperAdmin(dir *migrate.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.MySQL, w)), ent.Debug())
	ctx := context.Background()

	// The statement that generates the INSERT statement.
	err := client.Admin.Create().
		SetDisplayName("ADMIN").
		SetName("admin").AddRoles(client.Role.Create().SetName("Super Admin").SaveX(ctx)).
		Exec(context.Background())

	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"seed_admin",
		"Add the initial admin to the database.",
	)
}
