package interfaces

import "context"

type MessageLocationOutputDTO struct {
	Location string
	Message  string
}

type MessageRepository interface {
	ListByLocations(ctx context.Context, locations []string) (map[string]string, error)
}
