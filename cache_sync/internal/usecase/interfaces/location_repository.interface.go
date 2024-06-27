package interfaces

import (
	"cache_sync/internal/entity"
	"context"
)

type LocationRepository interface {
	ListUniqueLocations(ctx context.Context, page, size int) ([]entity.Location, error)
}
