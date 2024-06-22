package interfaces

import "context"

type ListActivesOutputDTO struct {
	Page    int
	Size    int
	Content []struct {
		ID    string
		Email string
		Phone string
	}
}

type UserRepository interface {
	save(ctx context.Context, id string, email string, phone string, active bool) error
	listActives(ctx context.Context, page, size int) (*ListActivesOutputDTO, error)
}
