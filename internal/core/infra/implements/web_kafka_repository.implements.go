package implements

import (
	"context"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
)

type WebKafkaRepository struct{}

func (k *WebKafkaRepository) Send(ctx context.Context, notification *entity.Notification) error {
	//TODO implement me
	fmt.Println("send notification to kafka queue", notification)
	return nil
}
