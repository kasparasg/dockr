package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kasparasg/dockr/api"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.NewApi()
}
