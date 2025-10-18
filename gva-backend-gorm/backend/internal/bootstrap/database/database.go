package database

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"backend/env"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type TxOperaton func(tx *gorm.DB) error

type Database struct {
	*gorm.DB
	cfg *env.Config
	log *zap.Logger
}

type Seeder interface {
	Name() string
	Count(ctx context.Context, conn *gorm.DB) (int, error)
	Seed(ctx context.Context, conn *gorm.DB) error
}

func NewDatabase(cfg *env.Config, log *zap.Logger) *Database {
	db := &Database{
		cfg: cfg,
		log: log,
		DB:  new(gorm.DB),
	}

	return db
}

func (db *Database) Connect() error {
	drv, err := sql.Open("mysql", db.cfg.DB.Mysql.DSN)
	if err != nil {
		return fmt.Errorf("dns %sv, An unknown error occurred when to connect the database!, %v", db.cfg.DB.Mysql.DSN, err)
	}

	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: drv,
	}), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return err
	}

	*db.DB = *gormdb
	db.log.Info("database is connected")
	return nil
}

func (db *Database) Close() error {
	defer db.log.Info("Database connection is closed")
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("an unknown error occurred when to shutdown the database! %v", err)
	}

	return nil
}

func (db *Database) MultiTransaction(fns ...TxOperaton) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, fn := range fns {
			if err := fn(db.DB); err != nil {
				return err
			}
		}

		return nil
	})
}

func (db *Database) SeedModels(ctx context.Context, seeder ...Seeder) {
	if !db.cfg.Seed.Enable {
		return
	}

	for _, v := range seeder {
		seedType := reflect.TypeOf(v).Elem().Name()
		seedlog := db.log.With(
			zap.String("name", v.Name()),
			zap.String("type", seedType),
		)

		if slices.ContainsFunc(db.cfg.Seed.BlacklistTypes, func(t string) bool {
			return strings.Contains((strings.ToLower(t)), strings.ToLower(seedType))
		}) {
			seedlog.Info("in blacklist. Skipping!")
			continue
		}

		count, err := v.Count(ctx, db.DB)

		if err != nil {
			seedlog.Error("v.Count(ctx, db.Client)", zap.Error(err))
			continue
		}

		if count > 0 {
			seedlog.Warn("Table has seeded already. Skipping!")
			continue
		}

		if err = v.Seed(ctx, db.DB); err != nil {
			seedlog.Error("Error seeding data", zap.Error(err))
			continue
		}

		seedlog.Info("Table has seeded successfully.")
	}

	db.log.Info("Seeding was completed!")
}
