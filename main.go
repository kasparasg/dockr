package main

import (
	"bytes"
	"log"

	"github.com/fsouza/go-dockerclient"
	"github.com/joho/godotenv"
)

func main() {
	var output bytes.Buffer

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, _ := docker.NewClientFromEnv()

	opts := docker.BuildImageOptions{
		Name:           "kasparasg/foo",
		SuppressOutput: true,
		OutputStream:   &output,
		Remote:         "express.tar",
	}
	err = client.BuildImage(opts)

	if err != nil {
		log.Fatal(err)
	}
}
