package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"time"
)

type ListNotExecutedByIntervalOutputDTO struct {
	Page    int
	Size    int
	Content []entity.ScheduleNotification
}

type ScheduleNotificationRepository interface {
	Save(ctx context.Context, scheduleNotification *entity.ScheduleNotification) error
	Delete(ctx context.Context, id string) error
	ListNotExecutedByInterval(ctx context.Context, page, size int, start time.Time, end time.Time) (*ListNotExecutedByIntervalOutputDTO, error)
}
