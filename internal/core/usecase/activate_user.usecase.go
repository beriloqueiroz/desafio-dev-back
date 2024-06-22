package usecase

import "context"

type ActivateUserUseCase struct {
	UserRepository interface{}
}

func (u *ActivateUserUseCase) Execute(ctx context.Context, ID string) error {
	return nil
}
