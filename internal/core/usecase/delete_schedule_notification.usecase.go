package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
)

type DeleteScheduleNotificationUseCase struct {
	ScheduleRepository interfaces.ScheduleNotificationRepository
}

func (u *DeleteScheduleNotificationUseCase) Execute(ctx context.Context, ID string) error {
	err := u.ScheduleRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
