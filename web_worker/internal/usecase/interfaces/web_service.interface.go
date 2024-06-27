package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/web_worker/internal/entity"
)

type WebService interface {
	Send(ctx context.Context, notifications []entity.Notification) error
}
