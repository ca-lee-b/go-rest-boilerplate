package main

import (
	"test/internal/api"
	"test/internal/api/handlers"
	"test/internal/log"
	"test/internal/repository"
)

func main() {
	log := log.New()

	repositories, err := repository.New()
	if err != nil {
		log.Error("Failed to connect to database")
	}

	handlers := handlers.New(repositories, log)
	api := api.New(handlers, log)

	err = api.Listen(8080)
	if err != nil {
		log.Error("Failed to start server: %v", err)
	}
}
