package interfaces

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type ListActivesOutputDTO struct {
	Page    int
	Size    int
	Content []entity.User
}

type UserRepository interface {
	Find(ctx context.Context, id string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	ListActives(ctx context.Context, page, size int) (*ListActivesOutputDTO, error)
}
