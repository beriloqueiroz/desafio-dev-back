package interfaces

import (
	"context"
	"core/internal/entity"
	"time"
)

type ScheduleNotificationRepository interface {
	Save(ctx context.Context, scheduleNotification *entity.ScheduleNotification) error
	Delete(ctx context.Context, id string) error
	FindFirstPendingBeforeDate(ctx context.Context, date time.Time) (*entity.ScheduleNotification, error)
}
