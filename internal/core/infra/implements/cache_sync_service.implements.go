package implements

import (
	"context"
)

type CacheSyncService struct{}

func (c *CacheSyncService) ListByLocations(ctx context.Context, locations []string) (map[string]string, error) {
	//TODO implement me
	panic("implement me")
}
