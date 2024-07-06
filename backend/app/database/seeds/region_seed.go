package seeds

import (
	"context"
	"log"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/ent"
	"github.com/gva/internal/ent/region"
	"github.com/gva/utils/json"
)

var _ interface{ database.Seeder } = (*RegionSeeder)(nil)

type RegionSeeder struct {
}

func (RegionSeeder) Count(ctx context.Context, db *ent.Client) (int, error) {
	return db.Region.Query().Count(ctx)
}

func (RegionSeeder) Seed(ctx context.Context, db *ent.Client) error {
	creates := make([]*ent.RegionCreate, 0)

	hasInsert := map[string]struct{}{}
	for _, c := range getCountryData() {
		if _, ok := hasInsert[c.NameID]; ok {
			continue
		}

		hasInsert[c.NameID] = struct{}{}
		creates = append(creates,
			db.Region.Create().
				SetName(c.Name).
				SetNameID(c.NameID).
				SetType(region.TypeCountry),
		)
	}

	return database.WithTx(ctx, db, func(tx *ent.Tx) error {
		_, err := tx.Region.CreateBulk(creates...).Save(ctx)
		return err
	})
}

func getCountryData() (countries []*ent.Region) {
	bytes, err := json.ReadJsonFile("./app/database/data/country_data.json")
	if err != nil {
		log.Panicf("can't raed seed data %v", err)
	}

	if err := bytes.Out(&countries); err != nil {
		log.Panicf("can't parse seed data %v", err)
	}

	return countries
}
