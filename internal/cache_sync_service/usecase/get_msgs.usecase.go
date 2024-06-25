package usecase

import (
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase/interfaces"
)

type GetMsgsUseCase struct {
	CacheRepository interfaces.CacheRepository
	MessageGateway  interfaces.MessageGateway
}

func NewGetMsgsUseCase(cacheRepository interfaces.CacheRepository) *GetMsgsUseCase {
	return &GetMsgsUseCase{
		CacheRepository: cacheRepository,
	}
}

func (u *GetMsgsUseCase) Execute(locations []entity.Location) (map[string]string, error) {
	// captura no cache e retorna
	// se não tem, captura do message gateway e popula cache
	panic("implement me")
}
