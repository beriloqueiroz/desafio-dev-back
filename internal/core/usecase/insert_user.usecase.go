package usecase

type InsertUserUseCase struct {
	UserRepository interface{}
}

type InsertUserUseCaseInput struct {
	Email string
	Phone string
}

func (u *InsertUserUseCase) Execute(input InsertUserUseCaseInput) error {
	return nil
}
