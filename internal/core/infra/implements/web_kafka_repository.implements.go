package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type WebKafkaRepository struct{}

func (k *WebKafkaRepository) Send(ctx context.Context, notification *entity.Notification) error {
	//TODO implement me
	return nil
}
