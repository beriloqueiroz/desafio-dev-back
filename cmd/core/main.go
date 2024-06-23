package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/infra/web"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	_ "github.com/lib/pq"
	"log"
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

	// repositories
	userRepository := implements.PostgresUserRepository{
		Db: db,
	}
	scheduleRepository := implements.PostgresScheduleRepository{}
	messageRepository := implements.CacheSyncService{}
	notificationQueueRepositories := []interfaces.NotificationQueueRepository{
		&implements.WebKafkaRepository{},
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
	webserver.AddRoute("POST /schedule-notification", scheduleRoutes.CreateScheduleNotificationHandler)
	webserver.AddRoute("DELETE /schedule-notification/{id}", scheduleRoutes.DeleteScheduleNotificationHandler)

	// start server
	srvErr := make(chan error, 1)
	go func() {
		fmt.Println("Starting web server on port", port)
		srvErr <- webserver.Start()
	}()

	// jobs
	go func() {
		for {
			fmt.Println("Starting sync schedules")
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
