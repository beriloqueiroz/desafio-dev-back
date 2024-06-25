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

func (u *InsertUserUseCase) Execute(ctx context.Context, email string, phone string, location string) (string, error) {
	loc, err := entity.LocationByString(location)
	if err != nil {
		return "", err
	}
	user, err := entity.NewUser(uuid.NewString(), true, email, phone, loc)
	if err != nil {
		return "", err
	}
	err = u.UserRepository.Save(ctx, user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
