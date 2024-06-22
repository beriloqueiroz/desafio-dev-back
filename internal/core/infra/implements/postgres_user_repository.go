package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type PostgresUserRepository struct{}

func (p *PostgresUserRepository) Find(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresUserRepository) Save(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresUserRepository) ListActives(ctx context.Context, page, size int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
