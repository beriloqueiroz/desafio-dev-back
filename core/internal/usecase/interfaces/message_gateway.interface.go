package interfaces

import (
	"context"
)

type MessageLocationOutputDTO struct {
	Location string
	Message  string
}

type MessageGateway interface {
	ListByLocations(ctx context.Context, locations []struct{ City, State string }) (map[string]string, error)
}
