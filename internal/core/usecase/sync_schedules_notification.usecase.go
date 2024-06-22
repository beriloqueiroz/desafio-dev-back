package usecase

import (
	"context"
	"time"
)

type SyncSchedulesNotificationUseCase struct {
	UserRepository     interface{}
	ScheduleRepository interface{}
	NotificationQueue  interface{}
}

func (u *SyncSchedulesNotificationUseCase) Execute(ctx context.Context, Message string, StartTime time.Time) error {
	return nil
}
