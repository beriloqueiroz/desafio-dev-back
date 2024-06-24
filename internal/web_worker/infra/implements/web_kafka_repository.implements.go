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
	msg_count := 0
	run := true
	min_commit_count := 2
	var notifications []entity.Notification
	for run == true {
		ev := k.Consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			msg_count += 1
			var notification entity.Notification
			err := json.Unmarshal(e.Value, &notification)
			if err != nil {
				slog.Error(fmt.Sprintf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value)))
				break
			}
			notifications = append(notifications, notification)
			ch <- notifications
			if msg_count%min_commit_count == 0 {
				go func() {
					_, err := k.Consumer.Commit()
					if err != nil {
						slog.Error(err.Error())
					}
				}()
			}
			slog.Info(fmt.Sprintf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value)))
		case kafka.PartitionEOF:
			msg := fmt.Sprintf("%% Reached %v\n", e)
			slog.Info(msg)
			run = false
			return nil
		case kafka.Error:
			msg := fmt.Sprintf("%% Error: %v\n", e)
			slog.Info(msg)
			run = false
			return errors.New(msg)
		default:
			slog.Info(fmt.Sprintf("Ignored %v\n", e))
			//return nil
		}
	}
	return nil
}
