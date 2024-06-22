package usecase

import "context"

type DeleteScheduleNotificationUseCase struct {
	ScheduleRepository interface{}
}

func (u *DeleteScheduleNotificationUseCase) Execute(ctx context.Context, ID string) error {
	return nil
}
