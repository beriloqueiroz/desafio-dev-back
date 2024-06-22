package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	"github.com/google/uuid"
)

type InsertUserUseCase struct {
	UserRepository interfaces.UserRepository
}

func (u *InsertUserUseCase) Execute(ctx context.Context, email string, phone string) error {
	user, err := entity.NewUser(uuid.NewString(), true, email, phone)
	if err != nil {
		return err
	}
	err = u.UserRepository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
