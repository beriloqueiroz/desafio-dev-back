package interfaces

import "context"

type CacheRepository interface {
	ListByKeys(ctx context.Context, keys []string) (map[string]string, error)
	FindByKey(ctx context.Context, key string) (string, error)
	Save(ctx context.Context, key string, value string) error
}
