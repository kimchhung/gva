package database

import (
	"context"
	dsql "database/sql"
	"fmt"
	"reflect"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gva/env"
	"github.com/gva/internal/ent"
	_ "github.com/gva/internal/ent/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Database struct {
	*ent.Client
	sql *sql.Driver
	Log *zerolog.Logger
	Cfg *env.Config
}

type Seeder interface {
	Count(ctx context.Context, conn *ent.Client) (int, error)
	Seed(ctx context.Context, conn *ent.Client) error
}

func NewDatabase(cfg *env.Config, log *zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Database) Connect() error {
	dbDns := fmt.Sprintf("%s?parseTime=true", db.Cfg.DB.Mysql.DSN)
	drv, err := sql.Open(dialect.MySQL, dbDns)
	if err != nil {
		return fmt.Errorf("dns %s, An unknown error occurred when to connect the database!, %v", db.Cfg.DB.Mysql.DSN, err)
	}

	db.sql = drv
	db.Client = ent.NewClient(
		ent.Driver(drv),
		ent.Debug(),
		ent.Log(log.Print),
	)

	if err := drv.DB().Ping(); err != nil {
		return fmt.Errorf("ping dns %s, An unknown error occurred when to connect the database!, %v", db.Cfg.DB.Mysql.DSN, err)
	}

	db.Log.Info().Msg("Database is connected")
	return nil
}

func (db *Database) Sql() *dsql.DB {
	return db.sql.DB()
}

func (db *Database) Close() error {
	defer db.Log.Info().Msg("Database connection is closed")

	if err := db.Client.Close(); err != nil {
		return fmt.Errorf("an unknown error occurred when to shutdown the database! %v", err)
	}

	return nil
}

func (db *Database) SeedModels(ctx context.Context, seeder ...Seeder) {
	if !db.Cfg.Seed.Enable {
		return
	}

	for _, v := range seeder {
		name := reflect.TypeOf(v).Elem().Name()
		count, err := v.Count(ctx, db.Client)
		if err != nil {
			db.Log.Panic().Err(err).Msg("v.Count(ctx, db.Client)")
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

	db.Log.Info().Msg("Seeding was completed!")
}

func WithTx(ctx context.Context, client *ent.Client, fns ...func(tx *ent.Tx) error) error {
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

	for _, fn := range fns {
		if err := fn(tx); err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
