package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	"time"
)

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interfaces.ScheduleNotificationRepository
}

func (u *CreateScheduleNotificationUseCase) Execute(ctx context.Context, Message string, StartTime time.Time) error {
	return nil
}
