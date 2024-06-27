package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/cache_sync_service/internal/entity"
)

type LocationRepository interface {
	ListUniqueLocations(ctx context.Context, page, size int) ([]entity.Location, error)
}
