package database

import (
	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/ent"

	"github.com/rs/zerolog"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Ent *ent.Client
	Log *zerolog.Logger
	Cfg *config.Config
}

type Seeder interface {
	Count(conn *ent.Client) (int, error)
	Seed(conn *ent.Client) error
}

func NewDatabase(cfg *config.Config, log *zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Database) ConnectDatabase() {
	drv, err := sql.Open(dialect.MySQL, db.Cfg.DB.Mysql.DSN)
	if err != nil {
		db.Log.Error().Err(err).Msg("An unknown error occurred when to connect the database!")
	}

	db.Ent = ent.NewClient(ent.Driver(drv))
}

func (db *Database) ShutdownDatabase() {
	if err := db.Ent.Close(); err != nil {
		db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database!")
	}
}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count(db.Ent)
		if err != nil {
			db.Log.Panic().Err(err).Msg("")
		}

		if count == 0 {
			err = v.Seed(db.Ent)
			if err != nil {
				db.Log.Panic().Err(err).Msg("")
			}

			db.Log.Debug().Msg("Table has seeded successfully.")
		} else {
			db.Log.Warn().Msg("Table has seeded already. Skipping!")
		}
	}

	db.Log.Info().Msg("Seeding was completed!")
}
