package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase/interfaces"
)

type SyncUseCase struct {
	LocationRepository interfaces.LocationRepository
	CacheRepository    interfaces.CacheRepository
}

func NewSyncUseCase(locationRepository interfaces.LocationRepository, cacheRepository interfaces.CacheRepository) *SyncUseCase {
	return &SyncUseCase{
		LocationRepository: locationRepository,
		CacheRepository:    cacheRepository,
	}
}

func (u *SyncUseCase) Execute(ctx context.Context) {
	// captura locations
	// captura mensagens
	// popula cache
}
