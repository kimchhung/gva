package seeds

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
)

type DepartmentSeeder struct {
}

func NewDepartSeeder() database.Seeder {
	return &DepartmentSeeder{}
}

func (DepartmentSeeder) Count(ctx context.Context, db *ent.Client) (int, error) {
	return db.Department.Query().Count(ctx)
}

func (DepartmentSeeder) Seed(ctx context.Context, db *ent.Client) error {
	creates := make([]*ent.DepartmentCreate, 0)

	creates = append(creates,
		db.Department.Create().
			SetName("Management").
			SetNameID("management"),

		db.Department.Create().
			SetName("Marketing").
			SetNameID("marketing"),

		db.Department.Create().
			SetName("IT").
			SetNameID("it"),

		db.Department.Create().
			SetName("Sales").
			SetNameID("sales"),
	)

	return database.WithTx(ctx, db, func(tx *ent.Tx) error {
		_, err := tx.Department.CreateBulk(creates...).Save(ctx)
		return err
	})
}
