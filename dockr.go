package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kasparasg/dockr/api"
	"github.com/kasparasg/dockr/queue"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	queue.StartDispatcher(2)

	api.NewApi()
}
