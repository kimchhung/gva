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

func createBulkPermissionDto(conn *ent.Client, keys ...PermissionKey) []*ent.PermissionCreate {
	bulks := make([]*ent.PermissionCreate, len(keys))

	for i, key := range keys {
		group, _, err := key.Value()
		if err != nil {
			panic(err)
		}

		bulks[i] = conn.Permission.Create().
			SetGroup(string(group)).
			SetKey(string(key)).
			SetName(key.Name()).
			SetOrder(i).
			SetType(permission.TypeStatic)
	}

	return bulks
}
