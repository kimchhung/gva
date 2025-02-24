package seeds

import (
	"backend/internal/bootstrap/database"
)

func AllSeeders() []database.Seeder {
	return []database.Seeder{
		NewSuperAdminSeeder(),
		NewConfigurationSeeder(),
	}
}
