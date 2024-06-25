package usecase

import "github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase/interfaces"

type SyncUseCase struct {
	LocationRepository interfaces.LocationRepository
	CacheRepository    interfaces.CacheRepository
}

func NewSyncUseCase() *SyncUseCase {
	return &SyncUseCase{}
}

func (u *SyncUseCase) Execute() {
	// captura locations
	// captura mensagens
	// popula cache
}
