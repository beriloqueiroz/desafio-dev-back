package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type NotificationQueueRepository interface {
	Send(ctx context.Context, notification *entity.Notification) error
}
