package interfaces

import (
	"context"
	"core/internal/entity"
)

type NotificationQueueRepository interface {
	Send(ctx context.Context, notification *entity.Notification) error
}
