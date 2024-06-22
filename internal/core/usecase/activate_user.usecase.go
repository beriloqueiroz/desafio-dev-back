package usecase

type ActivateUserUseCase struct {
	UserRepository interface{}
}

func (u *ActivateUserUseCase) Execute(ID string) error {
	return nil
}
