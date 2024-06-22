package interfaces

import (
	"context"
	"time"
)

type ListNotExecutedByIntervalOutputDTO struct {
	Page    int
	Size    int
	Content []struct {
		ID        string
		StartTime time.Time
		Executed  bool
	}
}

type ScheduleNotificationRepository interface {
	save(ctx context.Context, id string, email string, phone string, active bool) error
	delete(ctx context.Context, id string) error
	listNotExecutedByInterval(ctx context.Context, page, size int, start time.Time, end time.Time) ([]ListNotExecutedByIntervalOutputDTO, error)
}
