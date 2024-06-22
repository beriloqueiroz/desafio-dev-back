package usecase

import "context"

type InsertUserUseCase struct {
	UserRepository interface{}
}

func (u *InsertUserUseCase) Execute(ctx context.Context, Email string, Phone string) error {
	return nil
}
