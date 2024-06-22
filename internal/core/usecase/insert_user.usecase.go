package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
)

type InsertUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *InsertUserUseCase) Execute(ctx context.Context, Email string, Phone string) error {
	return nil
}
