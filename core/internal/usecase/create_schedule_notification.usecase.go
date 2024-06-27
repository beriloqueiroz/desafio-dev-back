package usecase

import (
	"context"
	"core/internal/entity"
	"core/internal/usecase/interfaces"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interfaces.ScheduleNotificationRepository
}

func (u *CreateScheduleNotificationUseCase) Execute(ctx context.Context, startTime time.Time) error {
	if startTime.Before(time.Now()) {
		return errors.New("start time is earlier than current time")
	}
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
