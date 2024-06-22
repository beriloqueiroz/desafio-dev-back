package usecase

import (
	"context"
	"time"
)

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interface{}
}

func (u *CreateScheduleNotificationUseCase) Execute(ctx context.Context, Message string, StartTime time.Time) error {
	return nil
}
