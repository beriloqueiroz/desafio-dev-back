package main

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"os/signal"
	"time"
)

func main() {
	// graceful exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	initCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// config logs
	// aqui no lugar do Stdout poderia ser um db ou elasticsearch por exemplo
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// load environment configs
	configs, err := configs.LoadConfig([]string{"."})
	if err != nil {
		panic(err)
	}

	kafkaNotificationQueueConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KAFKAUrl,
		"group.id":          "main",
		"auto.offset.reset": "smallest"})
	if err != nil {
		panic(err)
	}
	err = kafkaNotificationQueueConsumer.SubscribeTopics([]string{configs.KAFKATopic}, nil)
	if err != nil {
		panic(err)
	}
	defer kafkaNotificationQueueConsumer.Close()

	// repositories
	webKafkaRepository := implements.NewWebKafkaRepository(kafkaNotificationQueueConsumer, configs.KAFKATopic)
	webServiceClient := implements.NewWebRestService(configs.WebAppUrl)

	// useCases
	syncNotificationUseCase := usecase.SyncNotificationUseCase{
		WebService:         webServiceClient,
		NotificationQueues: webKafkaRepository,
	}

	// jobs
	go func() {
		slog.Info("Starting sync notifications")
		err := syncNotificationUseCase.Execute(context.Background())
		if err != nil {
			slog.Error(err.Error())
		}
		time.Sleep(time.Second)

	}()

	// Wait for interruption.
	select {
	case <-sigCh:
		slog.Warn("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		slog.Warn("Shutting down due to other reason...")
	}
}
