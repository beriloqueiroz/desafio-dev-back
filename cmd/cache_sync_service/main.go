package main

import (
	"context"
	"database/sql"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/infra/implements"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/usecase"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     configs.CacheUri,
		Password: configs.CachePass,
		DB:       0,
	})

	redisCacheRepository := implements.NewRedisCacheRepository(clientRedis)

	// repositories
	locationRepository := &implements.PostgresLocationRepository{
		Db: db,
	}

	syncUseCase := usecase.NewSyncUseCase(locationRepository, redisCacheRepository)

	go func() {
		for {
			syncUseCase.Execute()
			time.Sleep(time.Hour * 12) // o 12 pode ser variável de ambiente
		}
	}()

	// Wait for interruption.
	select {
	case <-sigCh:
		slog.Warn("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		slog.Warn("Shutting down due to other reason...")
	}
}
