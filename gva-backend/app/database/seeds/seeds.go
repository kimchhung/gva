package seeds

import (
	"github.com/gva/internal/bootstrap/database"
)

func AllSeeders() []database.Seeder {
	return []database.Seeder{
		NewDepartSeeder(),
		NewMenuSeeder(),
		NewMenuSeeder(),
		NewSuperAdminSeeder(),
	}
}
