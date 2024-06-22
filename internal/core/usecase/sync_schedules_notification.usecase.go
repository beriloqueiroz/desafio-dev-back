package usecase

import "time"

type SyncSchedulesNotificationUseCase struct {
	UserRepository     interface{}
	ScheduleRepository interface{}
	NotificationQueue  interface{}
}

type SyncSchedulesNotificationUseCaseInput struct {
	Message   string
	StartTime time.Time
}

func (u *SyncSchedulesNotificationUseCase) Execute(input SyncSchedulesNotificationUseCaseInput) error {
	return nil
}
