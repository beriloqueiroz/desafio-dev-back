package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	"github.com/google/uuid"
	"time"
)

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interfaces.ScheduleNotificationRepository
}

func (u *CreateScheduleNotificationUseCase) Execute(ctx context.Context, startTime time.Time) error {
	scheduleNotification, err := entity.NewScheduleNotification(uuid.NewString(), startTime, entity.Pending)
	if err != nil {
		return err
	}
	err = u.ScheduleRepository.Save(ctx, scheduleNotification)
	if err != nil {
		return err
	}
	return nil
}
