package main

import (
	"context"
	"database/sql"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/infra/web"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/lib/pq"
	"log"
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

	// dbs
	db, err := sql.Open(configs.DBDriver, configs.DBUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	kafkaNotificationQueueProduce, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KAFKAUrl,
	})
	if err != nil {
		panic(err)
	}
	defer kafkaNotificationQueueProduce.Close()

	// repositories
	userRepository := implements.PostgresUserRepository{
		Db: db,
	}
	scheduleRepository := implements.PostgresScheduleRepository{
		Db: db,
	}
	messageRepository := implements.CacheSyncService{}
	notificationQueueRepositories := []interfaces.NotificationQueueRepository{
		&implements.WebKafkaRepository{
			Producer: kafkaNotificationQueueProduce,
			Topic:    configs.KAFKATopic,
		},
	}

	// useCases
	insertUserUseCase := usecase.InsertUserUseCase{
		UserRepository: &userRepository,
	}
	activateUserRepository := usecase.ActivateUserUseCase{
		UserRepository: &userRepository,
	}
	deactivateUserRepository := usecase.DeactivateUserUseCase{
		UserRepository: &userRepository,
	}

	createScheduleNotificationUseCase := usecase.CreateScheduleNotificationUseCase{
		ScheduleRepository: &scheduleRepository,
	}
	deleteScheduleNotificationUseCase := usecase.DeleteScheduleNotificationUseCase{
		ScheduleRepository: &scheduleRepository,
	}
	syncSchedulesUseCase := usecase.SyncSchedulesNotificationUseCase{
		ScheduleRepository: &scheduleRepository,
		UserRepository:     &userRepository,
		MessageRepository:  &messageRepository,
		NotificationQueues: notificationQueueRepositories,
	}

	// webserver and routes
	port := ":8000"
	webserver := web.NewWebServer(port)

	userRoutes := web.NewUserRoutes(insertUserUseCase, activateUserRepository, deactivateUserRepository)
	webserver.AddRoute("POST /user", userRoutes.InsertUserHandler)
	webserver.AddRoute("PUT /user/{id}/activate", userRoutes.ActivateUserHandler)
	webserver.AddRoute("PUT /user/{id}/deactivate", userRoutes.DeactivateUserHandler)

	scheduleRoutes := web.NewSchedulerRoutes(createScheduleNotificationUseCase, deleteScheduleNotificationUseCase)
	webserver.AddRoute("POST /schedule", scheduleRoutes.CreateScheduleNotificationHandler)
	webserver.AddRoute("DELETE /schedule/{id}", scheduleRoutes.DeleteScheduleNotificationHandler)

	// start server
	srvErr := make(chan error, 1)
	go func() {
		slog.Info("Starting web server", "on port", port)
		srvErr <- webserver.Start()
	}()

	// jobs
	go func() {
		for {
			slog.Info("Starting sync schedules")
			err := syncSchedulesUseCase.Execute(context.Background())
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// Wait for interruption.
	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		log.Println("Shutting down due to other reason...")
	}
}
