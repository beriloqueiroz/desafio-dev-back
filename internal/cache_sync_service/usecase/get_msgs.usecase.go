package usecase

import (
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase/interfaces"
)

type GetMsgsUseCase struct {
	LocationRepository interfaces.LocationRepository
}

func NewGetMsgsUseCase(locationRepository interfaces.LocationRepository) *GetMsgsUseCase {
	return &GetMsgsUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *GetMsgsUseCase) Execute(locations []entity.Location) (map[string]string, error) {
	//TODO implement me
	res := make(map[string]string)
	for _, loc := range locations {
		res[loc.String()] = loc.String() + " teste calor quintura e mormaço"
	}
	return res, nil
}
