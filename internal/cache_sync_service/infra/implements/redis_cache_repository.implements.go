package implements

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheRepository struct {
	Client *redis.Client
}

func NewRedisCacheRepository(client *redis.Client) *RedisCacheRepository {
	return &RedisCacheRepository{
		Client: client,
	}
}

func (r *RedisCacheRepository) FindByKey(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCacheRepository) Save(ctx context.Context, key string, value string) error {
	err := r.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCacheRepository) SaveAll(ctx context.Context, values map[string]string) error {
	for k, v := range values {
		err := r.Client.HSet(ctx, time.Now().String(), k, v).Err()
		if err != nil {
			panic(err)
		}
	}
	return nil
}
