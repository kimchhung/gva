package database

import (
	"context"
	"fmt"
	"reflect"

	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/ent"

	"github.com/rs/zerolog"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*ent.Client
	Log *zerolog.Logger
	Cfg *config.Config
}

type Seeder interface {
	Count(ctx context.Context, conn *ent.Client) (int, error)
	Seed(ctx context.Context, conn *ent.Client) error
}

func NewDatabase(cfg *config.Config, log *zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Database) ConnectDatabase() {
	defer db.Log.Info().Msg("Database is connected!")

	drv, err := sql.Open(dialect.MySQL, db.Cfg.DB.Mysql.DSN)
	if err != nil {
		db.Log.Panic().Err(err).Str("dns", db.Cfg.DB.Mysql.DSN).Msg("An unknown error occurred when to connect the database!")
	}

	db.Client = ent.NewClient(ent.Driver(drv))
}

func (db *Database) ShutdownDatabase() {
	defer db.Log.Info().Msg("Database connection is closed")

	if err := db.Client.Close(); err != nil {
		db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database!")
	}
}

func (db *Database) SeedModels(ctx context.Context, seeder ...Seeder) {
	if !db.Cfg.Seed.Enable {
		return
	}

	defer db.Log.Info().Msg("Seeding was completed!")

	for _, v := range seeder {
		name := reflect.TypeOf(v).Name()
		count, err := v.Count(ctx, db.Client)
		if err != nil {
			db.Log.Panic().Err(err).Msg("")
		}

		if count > 0 {
			db.Log.Warn().Str("name", name).Msg("Table has seeded already. Skipping!")
			continue
		}

		if err = v.Seed(ctx, db.Client); err != nil {
			db.Log.Error().Err(err).Msg("Error seeding data")
			continue
		}

		db.Log.Debug().Str("name", name).Msg("Table has seeded successfully.")
	}
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
