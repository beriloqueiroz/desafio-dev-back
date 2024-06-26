package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase/interfaces"
	"log/slog"
)

type GetMsgsUseCase struct {
	CacheRepository interfaces.CacheRepository
	MessageGateway  interfaces.MessageGateway
}

func NewGetMsgsUseCase(cacheRepository interfaces.CacheRepository, messageGateway interfaces.MessageGateway) *GetMsgsUseCase {
	return &GetMsgsUseCase{
		CacheRepository: cacheRepository,
		MessageGateway:  messageGateway,
	}
}

func (u *GetMsgsUseCase) Execute(ctx context.Context, locations []entity.Location) (map[string]string, error) {
	// captura no cache e retorna
	result := make(map[string]string)
	for _, location := range locations {
		res, err := u.CacheRepository.Find(ctx, location.String())
		if err != nil {
			slog.Warn(err.Error())
			res, err = u.MessageGateway.MessageByLocation(ctx, location.City, location.State)
			if err != nil {
				slog.Error(err.Error())
				continue
			}
			err = u.CacheRepository.Save(ctx, location.String(), res)
			if err != nil {
				slog.Error(err.Error())
			}
			// se não tem, captura do message gateway e popula cache
		}
		result[location.String()] = res
	}
	return result, nil
}
