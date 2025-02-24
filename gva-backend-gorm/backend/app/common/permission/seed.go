package permission

import (
	"backend/app/common/model"
	"backend/internal/bootstrap/database"
	"time"
)

var (
	allSeeders []database.Seeder
)

func AllSeeders() []database.Seeder {
	return append([]database.Seeder{}, allSeeders...)
}

func createBulkPermissionDto(scopes ...permissionScope) []*model.Permission {
	bulks := make([]*model.Permission, len(scopes))

	for i, scope := range scopes {
		group, _, err := scope.Value()
		if err != nil {
			panic(err)
		}

		bulks[i] = &model.Permission{
			CreatedAt: time.Now(),
			Group:     string(group),
			Scope:     string(scope),
			Order:     i,
			Name:      scope.Name(),
		}
	}

	return bulks
}
