package infra

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheRepository struct {
	redisConn *redis.Client
}

func NewRedisCacheRepository(redisConn *redis.Client) *RedisCacheRepository {
	return &RedisCacheRepository{
		redisConn: redisConn,
	}
}

func (rcr *RedisCacheRepository) Get(ctx context.Context, key string) (interface{}, error) {
	res := rcr.redisConn.Get(ctx, key)
	if res.Err() == redis.Nil {
		return nil, errors.New("Key not exists")
	}
	return res.Val(), nil
}

func (rcr *RedisCacheRepository) SetIfNotExists(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ok, err := rcr.redisConn.SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		return fmt.Errorf("RedisCacheRepository[Set]: %w", err)
	}

	if ok {
		return nil
	}
	return errors.New("RedisCacheRepository[Set]: " + "Key already exists")
}