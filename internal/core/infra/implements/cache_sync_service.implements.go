package implements

import (
	"context"
)

type CacheSyncService struct{}

func (c *CacheSyncService) ListByLocations(ctx context.Context, locations []string) (map[string]string, error) {
	//TODO implement me
	res := make(map[string]string)
	for _, loc := range locations {
		res[loc] = loc + " teste calor quintura e mormaço"
	}
	return res, nil
}
