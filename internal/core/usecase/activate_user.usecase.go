package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
)

type ActivateUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *ActivateUserUseCase) Execute(ctx context.Context, ID string) error {
	return nil
}
