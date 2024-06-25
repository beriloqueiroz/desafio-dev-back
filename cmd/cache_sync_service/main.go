package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/configs"
	"github.com/beriloqueiroz/desafio-dev-back/internal/cache_sync_service/infra/implements"

	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"os/signal"
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

	// repositories
	locationRepository := implements.PostgresLocationRepository{
		Db: db,
	}

	res, err := locationRepository.ListUniqueLocations(context.Background(), 1, 500)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// Wait for interruption.
	select {
	case <-sigCh:
		slog.Warn("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		slog.Warn("Shutting down due to other reason...")
	}
}
