package interfaces

import "context"

type LocationRepository interface {
	ListUniquesLocations(ctx context.Context, page, size int) ([]string, error)
}
