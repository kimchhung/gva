package database

import (
	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/ent"
	"go.uber.org/zap"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Ent *ent.Client
	log *zap.Logger
	Cfg *config.Config
}

type Seeder interface {
	Count(conn *ent.Client) (int, error)
	Seed(conn *ent.Client) error
}

func NewDatabase(cfg *config.Config, log *zap.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		log: log,
	}

	return db
}

func (db *Database) ConnectDatabase() {
	drv, err := sql.Open(dialect.MySQL, db.Cfg.DB.Mysql.DSN)
	if err != nil {
		db.log.Error("An unknown error occurred when to connect the database!", zap.Error(err))
	}

	db.Ent = ent.NewClient(ent.Driver(drv))
}

func (db *Database) ShutdownDatabase() {
	if err := db.Ent.Close(); err != nil {
		db.log.Error("An unknown error occurred when to connect the database!", zap.Error(err))
	}
}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count(db.Ent)
		if err != nil {
			db.log.Panic("SeedModels", zap.Error(err))
		}

		if count == 0 {
			err = v.Seed(db.Ent)
			if err != nil {
				db.log.Panic("SeedModels", zap.Error(err))
			}

			db.log.Debug("Table has seeded successfully.", zap.Error(err))
		} else {
			db.log.Warn("Table has seeded already. Skipping!")
		}
	}

	db.log.Info("Seeding was completed!")
}
