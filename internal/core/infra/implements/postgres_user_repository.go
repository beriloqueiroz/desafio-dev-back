package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"github.com/google/uuid"
	"strconv"
)

type PostgresUserRepository struct{}

func (p *PostgresUserRepository) Find(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	return entity.NewUser(uuid.NewString(), true, "teste@teste.com", "12365478", "Fortaleza")
}

func (p *PostgresUserRepository) Save(ctx context.Context, user *entity.User) error {
	//TODO implement me
	return nil
}

func (p *PostgresUserRepository) ListActives(ctx context.Context, page, size int) ([]entity.User, error) {
	//TODO implement me
	var result []entity.User
	for i := 0; i < size; i++ {
		user, _ := entity.NewUser(uuid.NewString(), true, "teste@teste.com"+strconv.Itoa(i), "12365478"+strconv.Itoa(i), "Fortaleza")
		result = append(result, *user)
	}
	return result, nil
}
