package database

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"backend/core/env"
	coretype "backend/core/type"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xo/dburl"
)

type Database struct {
	*gorm.DB
	cfg *env.Config
	log *zap.Logger
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
	url, err := dburl.Parse(db.cfg.DB.Url)
	if err != nil {
		return fmt.Errorf("failed to parse db url: %v", err)
	}

	drv, err := sql.Open(url.Driver, url.DSN)
	if err != nil {
		return fmt.Errorf("dns %sv, An unknown error occurred when to connect the database!, %v", url.DSN, err)
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

func (db *Database) SeedModels(ctx context.Context, seeder ...coretype.Seeder) {
	if !db.cfg.DB.Seed.Enable {
		return
	}

	for _, v := range seeder {
		seedType := reflect.TypeOf(v).Elem().Name()
		seedlog := db.log.With(
			zap.String("name", v.Name()),
			zap.String("type", seedType),
		)

		if slices.ContainsFunc(db.cfg.DB.Seed.BlacklistTypes, func(t string) bool {
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
