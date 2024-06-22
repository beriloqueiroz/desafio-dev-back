package usecase

type DeactivateUserUseCase struct {
	UserRepository interface{}
}

type DeactivateUserUseCaseInput struct {
	ID string
}

func (u *DeactivateUserUseCase) Execute(input DeactivateUserUseCaseInput) error {
	return nil
}
