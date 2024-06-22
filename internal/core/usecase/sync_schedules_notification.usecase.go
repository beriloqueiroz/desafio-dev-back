package usecase

import "time"

type SyncSchedulesNotificationUseCase struct {
	UserRepository     interface{}
	ScheduleRepository interface{}
	NotificationQueue  interface{}
}

func (u *SyncSchedulesNotificationUseCase) Execute(Message string, StartTime time.Time) error {
	return nil
}
