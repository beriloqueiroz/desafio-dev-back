package usecase

import "time"

type CreateScheduleNotificationUseCase struct {
	ScheduleRepository interface{}
}

type CreateScheduleNotificationUseCaseInput struct {
	Message   string
	StartTime time.Time
}

func (u *CreateScheduleNotificationUseCase) Execute(input CreateScheduleNotificationUseCaseInput) error {
	return nil
}
