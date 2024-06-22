package interfaces

import "context"

type NotificationQueueRepository interface {
	send(ctx context.Context, msg string) error
}
