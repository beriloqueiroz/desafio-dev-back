package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
)

type DeactivateUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *DeactivateUserUseCase) Execute(ctx context.Context, ID string) error {
	return nil
}
