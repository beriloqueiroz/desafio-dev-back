package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/cache_sync_service/internal/entity"
	interfaces2 "github.com/beriloqueiroz/desafio-dev-back/cache_sync_service/internal/usecase/interfaces"
	"log/slog"
)

type GetMsgsUseCase struct {
	CacheRepository interfaces2.CacheRepository
	MessageGateway  interfaces2.MessageGateway
}

func NewGetMsgsUseCase(cacheRepository interfaces2.CacheRepository, messageGateway interfaces2.MessageGateway) *GetMsgsUseCase {
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
