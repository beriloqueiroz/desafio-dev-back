package usecase

type InsertUserUseCase struct {
	UserRepository interface{}
}

func (u *InsertUserUseCase) Execute(Email string, Phone string) error {
	return nil
}
