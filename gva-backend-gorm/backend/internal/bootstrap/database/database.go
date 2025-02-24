package database

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"backend/env"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type TxOperaton func(tx *gorm.DB) error

type Database struct {
	*gorm.DB
	Cfg *env.Config
	Log *zap.Logger
}

type Seeder interface {
	Name() string
	Count(ctx context.Context, conn *gorm.DB) (int, error)
	Seed(ctx context.Context, conn *gorm.DB) error
}

func NewDatabase(cfg *env.Config, log *zap.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
		DB:  new(gorm.DB),
	}

	return db
}

func (db *Database) Connect() error {
	drv, err := sql.Open("mysql", db.Cfg.DB.Mysql.DSN)
	if err != nil {
		return fmt.Errorf("dns %sv, An unknown error occurred when to connect the database!, %v", db.Cfg.DB.Mysql.DSN, err)
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
	db.Log.Info("database is connected")
	return nil
}

func (db *Database) Close() error {
	defer db.Log.Info("Database connection is closed")
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
	if !db.Cfg.Seed.Enable {
		return
	}

	for _, v := range seeder {
		seedType := reflect.TypeOf(v).Elem().Name()
		seedlog := db.Log.With(
			zap.String("name", v.Name()),
			zap.String("type", seedType),
		)

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

	db.Log.Info("Seeding was completed!")
}
