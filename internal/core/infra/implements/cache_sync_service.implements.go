package implements

import (
	"context"
	cacheentity "github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/entity"
	cacheusecase "github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type CacheSyncService struct {
	MessageCacheUseCase *cacheusecase.GetMsgsUseCase
}

func NewCacheSyncService(messageCacheUseCase *cacheusecase.GetMsgsUseCase) *CacheSyncService {
	return &CacheSyncService{
		MessageCacheUseCase: messageCacheUseCase,
	}
}

func (c *CacheSyncService) ListByLocations(ctx context.Context, locations []entity.Location) (map[string]string, error) {
	var locs []cacheentity.Location
	for _, location := range locations {
		locs = append(locs, cacheentity.Location{
			location.City, location.State,
		})
	}
	return c.MessageCacheUseCase.Execute(ctx, locs)
}
