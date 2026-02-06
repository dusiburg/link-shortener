package main

import (
	"fmt"
	"log"
	"os"

	"example.com/m/internal/config"
	"example.com/m/internal/db"
	"example.com/m/internal/handlers"
	httpserver "example.com/m/internal/http"
	"example.com/m/internal/repositories"
	"example.com/m/internal/services"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("server has been stopped: %v", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configs: %v", err)
	}

	database, err := db.InitRedis(cfg)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	defer database.Close()

	address := fmt.Sprintf("0.0.0.0:%s", cfg.ServerPort)

	repo := repositories.NewLinksRepository(database)
	service := services.NewLinksService(repo)
	handler := handlers.NewLinksHandler(service, cfg.ServerHost, cfg.ServerPort)

	server := httpserver.NewServer(address, handler)
	return server.Start()
}
