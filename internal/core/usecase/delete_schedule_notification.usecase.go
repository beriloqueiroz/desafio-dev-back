package usecase

type DeleteScheduleNotificationUseCase struct {
	ScheduleRepository interface{}
}

func (u *DeleteScheduleNotificationUseCase) Execute(ID string) error {
	return nil
}
