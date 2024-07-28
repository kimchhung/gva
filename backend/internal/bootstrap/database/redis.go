package database

import (
	"context"
	"fmt"

	"github.com/gva/env"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Redis struct {
	*redis.Client
	Log *zerolog.Logger
	Cfg *env.Config
}

func NewRedis(cfg *env.Config, log *zerolog.Logger) *Redis {
	db := &Redis{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Redis) Connect() error {
	db.Client = redis.NewClient(&redis.Options{
		Addr:     db.Cfg.DB.Redis.Addr,
		Password: db.Cfg.DB.Redis.Password,
	})

	_, err := db.Client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	db.Log.Info().Msg("Redis is connected")
	return nil
}

func (db *Redis) Close() error {
	if err := db.Client.Close(); err != nil {
		return fmt.Errorf("failed to shutdown Redis: %v", err)
	}
	db.Log.Info().Msg("Redis connection is closed")
	return nil
}
