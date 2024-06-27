package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type MessageLocationOutputDTO struct {
	Location string
	Message  string
}

type MessageGateway interface {
	ListByLocations(ctx context.Context, locations []entity.Location) (map[string]string, error)
}
