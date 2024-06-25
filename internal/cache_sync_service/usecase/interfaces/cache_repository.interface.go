package interfaces

import "context"

type CacheRepository interface {
	FindByKey(ctx context.Context, key string) (string, error)
	Save(ctx context.Context, key string, value string) error
	SaveAll(ctx context.Context, values map[string]string) error
}
