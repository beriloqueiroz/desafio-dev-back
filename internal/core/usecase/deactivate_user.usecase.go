package usecase

type DeactivateUserUseCase struct {
	UserRepository interface{}
}

func (u *DeactivateUserUseCase) Execute(ID string) error {
	return nil
}
