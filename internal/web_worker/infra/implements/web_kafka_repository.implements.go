package implements

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
)

type WebKafkaRepository struct {
	Consumer     *kafka.Consumer
	Topic        string
	DeliveryChan chan kafka.Event
}

func NewWebKafkaRepository(consumer *kafka.Consumer, topic string) *WebKafkaRepository {
	deliveryChan := make(chan kafka.Event, 1000)
	go func() {
		for e := range deliveryChan {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					slog.Error(fmt.Sprintf("Failed to deliver message: %v\n", ev.TopicPartition))
				} else {
					slog.Info(fmt.Sprintf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset))
				}
			}
		}
	}()
	return &WebKafkaRepository{
		Consumer:     consumer,
		Topic:        topic,
		DeliveryChan: deliveryChan,
	}
}

func (k *WebKafkaRepository) Read(ctx context.Context, ch chan []entity.Notification) error {
	run := true
	err := k.Consumer.SubscribeTopics([]string{k.Topic}, nil)
	if err != nil {
		return err
	}
	defer k.Consumer.Close()
	var notifications []entity.Notification
	for run == true {
		ev := k.Consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			var notification entity.Notification
			err := json.Unmarshal(e.Value, &notification)
			if err != nil {
				slog.Error(fmt.Sprintf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value)))
				break
			}
			notifications = append(notifications, notification)
			//slog.Info(fmt.Sprintf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value)))
		case kafka.PartitionEOF:
			msg := fmt.Sprintf("%% Reached %v\n", e)
			slog.Info(msg)
			run = false
			return nil
		case kafka.Error:
			msg := fmt.Sprintf("%% Error: %v\n", e)
			run = false
			return errors.New(msg)
		default:
			if len(notifications) > 0 {
				ch <- notifications
				notifications = notifications[:0]
			}
		}
	}
	return nil
}

func (k *WebKafkaRepository) Commit(ctx context.Context) error {
	info, err := k.Consumer.Commit()
	slog.Info(fmt.Sprintf("Commited Topic %s offset %s \n", *info[len(info)-1].Topic, info[len(info)-1].Offset.String()))
	if err != nil {
		return err
	}
	return nil
}
