package main

import (
	"context"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/infra/web"
	"log"
	"os"
	"os/signal"
)

func main() {
	// graceful exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	initCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// load environment configs

	// servers and jobs
	port := ":8080"
	webserver := web.NewWebServer(port)
	userRoutes := web.NewUserRoutes()
	webserver.AddRoute("POST /user", userRoutes.CreateUserHandler)
	webserver.AddRoute("PUT /user/{id}/{active}", userRoutes.ActivateUserHandler)
	srvErr := make(chan error, 1)
	go func() {
		fmt.Println("Starting web server on port", port)
		srvErr <- webserver.Start()
	}()

	// Wait for interruption.
	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-initCtx.Done():
		log.Println("Shutting down due to other reason...")
	}
}
