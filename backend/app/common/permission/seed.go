package permission

import "github.com/gva/internal/bootstrap/database"

func AllSeeders() []database.Seeder {
	return []database.Seeder{
		NewAdminSeeder(),
		NewAdminRoleSeeder(),
		NewRouteSeeder(),
	}
}
