package main

import (
	"github.com/ca-lee-b/go-rest-boilerplate/internal/api"
	"github.com/ca-lee-b/go-rest-boilerplate/internal/api/handlers"
	"github.com/ca-lee-b/go-rest-boilerplate/internal/log"
	"github.com/ca-lee-b/go-rest-boilerplate/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Could not load env file")
	}

	logger := log.New()

	repositories, err := repository.New()
	if err != nil {
		logger.Error("Failed to connect to database")
	}

	handlers := handlers.New(repositories, logger)
	api := api.New(handlers, logger)

	err = api.Listen()
	if err != nil {
		logger.Error("Failed to start server: %v", err)
	}
}
