package seeds

import coretype "backend/core/type"

func AllSeeders() []coretype.Seeder {
	return []coretype.Seeder{
		NewSuperAdminSeeder(),
		NewConfigurationSeeder(),
	}
}
