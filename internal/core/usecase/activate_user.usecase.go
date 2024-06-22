package usecase

type ActivateUserUseCase struct {
	UserRepository interface{}
}

type ActivateUserUseCaseInput struct {
	ID string
}

func (u *ActivateUserUseCase) Execute(input ActivateUserUseCaseInput) error {
	return nil
}
