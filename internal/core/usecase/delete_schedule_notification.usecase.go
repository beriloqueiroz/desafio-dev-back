package usecase

type DeleteScheduleNotificationUseCase struct {
	ScheduleRepository interface{}
}

type DeleteScheduleNotificationUseCaseInput struct {
	ID string
}

func (u *DeleteScheduleNotificationUseCase) Execute(input DeleteScheduleNotificationUseCaseInput) error {
	return nil
}
