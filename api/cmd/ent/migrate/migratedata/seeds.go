package migratedata

import (
	"log"

	atlas "ariga.io/atlas/sql/migrate"
)

type SeedFunc = func(dir *atlas.LocalDir) error

var Seeds = []SeedFunc{
	SeedSuperAdmin,
}

func GenerateSql(dir *atlas.LocalDir) {
	for _, sf := range Seeds {
		if err := sf(dir); err != nil {
			log.Fatalf("failed to generate seed")
		}
	}
}
