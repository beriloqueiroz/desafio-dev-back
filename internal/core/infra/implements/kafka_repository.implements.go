package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type KafkaRepository struct{}

func (k *KafkaRepository) Send(ctx context.Context, notification *entity.Notification) error {
	//TODO implement me
	return nil
}
