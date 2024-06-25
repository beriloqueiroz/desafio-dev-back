package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/entity"
)

type LocationRepository interface {
	ListUniquesLocations(ctx context.Context, page, size int) ([]entity.Location, error)
}
