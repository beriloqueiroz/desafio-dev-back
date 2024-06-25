package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type CacheSyncService struct{}

func (c *CacheSyncService) ListByLocations(ctx context.Context, locations []entity.Location) (map[string]string, error) {
	//TODO implement me
	res := make(map[string]string)
	for _, loc := range locations {
		res[loc.String()] = loc.String() + " teste calor quintura e mormaço"
	}
	return res, nil
}
