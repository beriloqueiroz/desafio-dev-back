package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
)

type Action func(ctx context.Context, notifications []entity.Notification) error

type NotificationQueueRepository interface {
	Read(ctx context.Context, action Action) error
}
