package web_worker

import (
	"context"
	"database/sql"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/infra/web"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/usecase"
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

	kafkaNotificationQueueConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KAFKAUrl})
	if err != nil {
		panic(err)
	}
	err = kafkaNotificationQueueConsumer.SubscribeTopics([]string{configs.KAFKATopic}, nil)
	if err != nil {
		panic(err)
	}
	defer kafkaNotificationQueueConsumer.Close()

	// repositories
	userRepository := implements.PostgresUserRepository{
		Db: db,
	}

	webKafkaRepository := implements.NewWebKafkaRepository(kafkaNotificationQueueConsumer, configs.KAFKATopic)
	webServiceClient := implements.NewWebRestService(configs.WebAppUrl)

	// useCases
	activateUserRepository := usecase.ActivateUserUseCase{
		UserRepository: &userRepository,
	}
	deactivateUserRepository := usecase.DeactivateUserUseCase{
		UserRepository: &userRepository,
	}
	syncNotificationUseCase := usecase.SyncNotificationUseCase{
		WebService:         webServiceClient,
		NotificationQueues: webKafkaRepository,
	}

	// webserver and routes
	port := ":3333"
	webserver := web.NewWebServer(port)

	userRoutes := web.NewUserRoutes(activateUserRepository, deactivateUserRepository)
	webserver.AddRoute("PUT /user/{id}/activate", userRoutes.ActivateUserHandler)
	webserver.AddRoute("PUT /user/{id}/deactivate", userRoutes.DeactivateUserHandler)

	// jobs
	go func() {
		for {
			slog.Info("Starting sync notifications")
			err := syncNotificationUseCase.Execute(context.Background())
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
