package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
)

type WebService interface {
	Send(ctx context.Context, notifications []entity.Notification) error
}
