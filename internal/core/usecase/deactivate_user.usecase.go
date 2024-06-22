package usecase

import "context"

type DeactivateUserUseCase struct {
	UserRepository interface{}
}

func (u *DeactivateUserUseCase) Execute(ctx context.Context, ID string) error {
	return nil
}
