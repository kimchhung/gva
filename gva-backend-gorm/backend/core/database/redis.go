package database

import (
	"backend/core/env"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	*redis.Client
	Log *zap.Logger
	Cfg *env.Config
}

func NewRedis(cfg *env.Config, log *zap.Logger) *Redis {
	db := &Redis{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Redis) Connect() error {
	if !(db.Cfg.DB.Redis.Enable) {
		return nil
	}

	url, err := redis.ParseURL(db.Cfg.DB.Redis.Url)
	if err != nil {
		return fmt.Errorf("failed to parse redis url: %v", err)
	}

	db.Client = redis.NewClient(url)
	_, err = db.Client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	db.Log.Info("redis is connected")
	return nil
}

func (db *Redis) Close() error {
	if err := db.Client.Close(); err != nil {
		return fmt.Errorf("failed to shutdown Redis: %v", err)
	}

	return nil
}
