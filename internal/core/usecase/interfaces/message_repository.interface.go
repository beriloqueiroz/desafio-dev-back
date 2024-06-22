package interfaces

import "context"

type MessageRepository interface {
	findByLocation(ctx context.Context, location string) (string, error)
}
