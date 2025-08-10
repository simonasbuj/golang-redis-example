package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-redis/redis/v8"
)


type RedisCache struct {
	client	*redis.Client
}


func NewRedisCache(opt *redis.Options) (*RedisCache, error) {
	slog.Info("creating redis client...")

	client := redis.NewClient(opt)

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	slog.Info("ping successful", "result", ping)

	return &RedisCache{
		client: client,
	}, nil
}

func (rc *RedisCache) SetValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := rc.client.Set(ctx, key, value, expiration).Err()
	if err != nil{
		slog.Error("failed to set value in the redis instance", "err", err.Error())
		return err
	}

	slog.Info(
		"value put into redis", 
		"key", key,
		"value", value,
		"expiration", expiration,
	)
	return nil
}

func (rc *RedisCache) GetValue(ctx context.Context, key string) (string, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		slog.Error("failed to get value from redis", "err", err.Error())
		return "", err
	}

	return val, nil
}