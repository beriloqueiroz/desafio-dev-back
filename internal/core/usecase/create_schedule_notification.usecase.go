package usecase

import "time"

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interface{}
}

func (u *CreateScheduleNotificationUseCase) Execute(Message string, StartTime time.Time) error {
	return nil
}
