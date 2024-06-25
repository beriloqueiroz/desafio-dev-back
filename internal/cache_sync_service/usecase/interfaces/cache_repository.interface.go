package interfaces

import "context"

type CacheRepository interface {
	ListByKeys(ctx context.Context, keys []string) (map[string]string, error)
}
