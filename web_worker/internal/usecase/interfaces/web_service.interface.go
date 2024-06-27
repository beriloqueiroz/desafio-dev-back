package interfaces

import (
	"context"
	"web_worker/internal/entity"
)

type WebService interface {
	Send(ctx context.Context, notifications []entity.Notification) error
}
