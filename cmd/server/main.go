package main

import (
	"real-time/internal/domains/server"
	"real-time/internal/handlers"
	"real-time/pkg/logger"
)

func main() {
	logger := logger.InitLogger()
	handler := handlers.NewHandler(logger)
	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logger.Fatalf("error occured running server: %s", err.Error())
	}
}
