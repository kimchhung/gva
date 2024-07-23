package permission

import (
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/permission"
)

var (
	allSeeders []database.Seeder
)

func AllSeeders() []database.Seeder {
	return append([]database.Seeder{}, allSeeders...)
}

func createBulkPermissionDto(conn *ent.Client, scopes ...permissionScope) []*ent.PermissionCreate {
	bulks := make([]*ent.PermissionCreate, len(scopes))

	for i, scope := range scopes {
		group, _, err := scope.Value()
		if err != nil {
			panic(err)
		}

		bulks[i] = conn.Permission.Create().
			SetGroup(string(group)).
			SetScope(string(scope)).
			SetName(scope.Name()).
			SetOrder(i).
			SetType(permission.TypeStatic)
	}

	return bulks
}
