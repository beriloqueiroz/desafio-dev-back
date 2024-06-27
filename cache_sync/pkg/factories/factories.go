package factories

import (
	"cache_sync/internal/entity"
	"cache_sync/internal/infra/implements"
	"cache_sync/internal/usecase"
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheSyncMessageGateway struct {
	cacheUri, cachePass string
}

func NewCacheSyncMessageGateway(cacheUri, cachePass string) *CacheSyncMessageGateway {
	return &CacheSyncMessageGateway{cacheUri, cachePass}
}

func (cs *CacheSyncMessageGateway) ListByLocations(ctx context.Context, locations []struct{ City, State string }) (map[string]string, error) {
	clientRedis := redis.NewClient(&redis.Options{
		Addr:     cs.cacheUri,
		Password: cs.cachePass,
		DB:       0,
	})
	redisCacheRepository := implements.NewRedisCacheRepository(clientRedis)
	cptecMessageGateway := implements.NewCptecMessageGateway()
	usecase := usecase.NewGetMsgsUseCase(redisCacheRepository, cptecMessageGateway)
	var locs []entity.Location
	for _, loc := range locations {
		locs = append(locs, entity.Location{
			City:  loc.City,
			State: loc.State,
		})
	}
	return usecase.Execute(ctx, locs)
}
