package implements

import (
	"context"
	"encoding/json"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
)

type WebKafkaRepository struct {
	Producer *kafka.Producer
	Topic    string
}

func (k *WebKafkaRepository) Send(ctx context.Context, notification *entity.Notification) error {
	value, err := json.Marshal(notification)
	if err != nil {
		slog.Error("Erro ao tentar serializar para kafka", "error", err)
		return err
	}
	err = k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.Topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
	if err != nil {
		slog.Error("Erro ao tentar enviar para kafka", "error", err)
		return err
	}
	return nil
}
