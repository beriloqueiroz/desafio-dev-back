package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/core/internal/entity"
)

type GetMsgUseCase interface {
	Execute(ctx context.Context, locations []entity.Location) (map[string]string, error)
}

type CacheSyncService struct {
	MessageCacheUseCase GetMsgUseCase
}

func NewCacheSyncService(messageCacheUseCase GetMsgUseCase) *CacheSyncService {
	return &CacheSyncService{
		MessageCacheUseCase: messageCacheUseCase,
	}
}

func (c *CacheSyncService) ListByLocations(ctx context.Context, locations []entity.Location) (map[string]string, error) {
	return c.MessageCacheUseCase.Execute(ctx, locations)
}
