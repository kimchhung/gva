package service

import (
	"backend/internal/bootstrap/database"
	"backend/utils/json"
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisService struct {
	redis *database.Redis
	log   *zap.Logger
}

func NewRedisService(redis *database.Redis, log *zap.Logger) *RedisService {
	return &RedisService{
		redis: redis,
		log:   log.Named("cacheService"),
	}
}

func (s *RedisService) Publish(ctx context.Context, channel string, value any) error {
	payload := json.JSON{}
	if err := payload.Set(value); err != nil {
		return err
	}

	if err := s.redis.Publish(ctx, channel, string(payload)).Err(); err != nil {
		return err
	}
	return nil
}

func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	rcmd := s.redis.Get(ctx, key)
	if err := rcmd.Err(); err != nil {
		if errors.Is(redis.Nil, err) {
			return "", nil
		}
		return "", err
	}
	return rcmd.String(), rcmd.Err()
}

// support primative type, except json
func (s *RedisService) GetTo(ctx context.Context, key string, value any) error {
	rcmd := s.redis.Get(ctx, key)
	if err := rcmd.Err(); err != nil {
		if errors.Is(redis.Nil, err) {
			return nil
		}
		return err
	}

	return rcmd.Scan(value)
}

func (s *RedisService) GetJsonTo(ctx context.Context, key string, value any) (err error) {
	rcmd := s.redis.Get(ctx, key)
	if err := rcmd.Err(); err != nil {
		if errors.Is(redis.Nil, err) {
			return nil
		}
		return err
	}

	var rawByte json.JSON
	rawByte, err = rcmd.Bytes()
	if err != nil {
		return err
	}

	if err = rawByte.Out(value); err != nil {
		return err
	}

	return nil
}

func (s *RedisService) WrapTo(ctx context.Context, key string, ttl time.Duration, to any, fn func() (any, error)) (err error) {
	rcmd := s.redis.Get(ctx, key)
	if err := rcmd.Err(); err == redis.Nil {
		fnValue, fnErr := fn()
		if fnErr != nil {
			return fnErr
		}

		if reflect.TypeOf(fnValue).AssignableTo(reflect.TypeOf(to).Elem()) {
			reflect.ValueOf(to).Elem().Set(reflect.ValueOf(fnValue))
		} else if reflect.ValueOf(fnValue).Kind() == reflect.Ptr &&
			reflect.TypeOf(fnValue).Elem().AssignableTo(reflect.TypeOf(to).Elem().Elem()) {
			reflect.ValueOf(to).Elem().Set(reflect.ValueOf(fnValue).Elem())
		} else {
			return fmt.Errorf("type not match %T %T", fnValue, to)
		}

		if err = s.SetJson(ctx, key, to, ttl); err != nil {
			return err
		}

	} else if err != nil {
		return err
	}

	var rawByte json.JSON
	rawByte, err = rcmd.Bytes()
	if err != nil {
		return err
	}

	if err = rawByte.Out(to); err != nil {
		return err
	}

	return nil
}

func (s *RedisService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	rcmd := s.redis.Set(ctx, key, value, ttl)
	return rcmd.Err()
}

func (s *RedisService) SetJson(ctx context.Context, key string, value any, ttl time.Duration) error {
	rcmd := s.redis.Set(ctx, key, json.MustJSON(value).String(), ttl)
	return rcmd.Err()
}

func (s *RedisService) HGet(ctx context.Context, key string, field string, result interface{}) error {
	var (
		bytes json.JSON
		err   error
	)

	rdc := s.redis.HGet(ctx, key, field)
	if rdc.Err() != nil {
		return err
	}

	bytes, err = rdc.Bytes()
	if err != nil {
		return err
	}

	if err = bytes.Out(&result); err != nil {
		return err
	}

	return err
}

func (s *RedisService) HVals(ctx context.Context, base string, result interface{}) error {
	rdc := s.redis.HVals(ctx, base)
	if err := rdc.Err(); err != nil {
		return err
	}

	listStr, err := rdc.Result()
	if err != nil {
		return err
	}

	reflectValue := reflect.ValueOf(result)
	if reflectValue.Kind() != reflect.Ptr || reflectValue.IsNil() {
		return errors.New("result must be a non-nil pointer")
	}

	sliceValue := reflectValue.Elem()
	if sliceValue.Kind() != reflect.Slice {
		return errors.New("result must be a pointer to a slice")
	}

	sliceValue.Set(reflect.MakeSlice(sliceValue.Type(), len(listStr), len(listStr)))
	elemType := sliceValue.Type().Elem()

	for i, item := range listStr {
		data := reflect.New(elemType).Interface()
		raw := json.JSON(item)
		if err := raw.Out(data); err != nil {
			return err
		}
		if raw.String() != "null" {
			sliceValue.Index(i).Set(reflect.ValueOf(data).Elem())
		}
	}

	return nil
}

func (s *RedisService) HDel(ctx context.Context, key string, fields ...string) error {
	return s.redis.HDel(ctx, key, fields...).Err()
}
func (s *RedisService) Del(ctx context.Context, keys ...string) (int64, error) {
	return s.Client().Del(ctx, keys...).Result()
}

func (s *RedisService) DeleteByPattern(ctx context.Context, pattern string) (int64, error) {
	keys, err := s.Client().Keys(ctx, pattern).Result()
	if err != nil {
		return 0, err
	}

	if len(keys) == 0 {
		return 0, nil
	}

	return s.Del(ctx, keys...)
}

/*
support key value

	HSet("myhash", "key1", "value1", "key2", "value2")
	HSet("myhash", []string{"key1", "value1", "key2", "value2"})
	HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
*/
func (s *RedisService) HSet(ctx context.Context, key string, values ...interface{}) error {
	return s.redis.HSet(ctx, key, values...).Err()
}

func (s *RedisService) Client() *redis.Client {
	return s.redis.Client
}
