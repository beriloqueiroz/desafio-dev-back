package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"time"
)

type ScheduleNotificationRepository interface {
	Save(ctx context.Context, scheduleNotification *entity.ScheduleNotification) error
	Delete(ctx context.Context, id string) error
	FindFirstNotExecutedBeforeDate(ctx context.Context, date time.Time) (entity.ScheduleNotification, error)
}
