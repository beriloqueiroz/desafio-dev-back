package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
)

type NotificationQueueRepository interface {
	Read(ctx context.Context, ch chan []entity.Notification) error
}
