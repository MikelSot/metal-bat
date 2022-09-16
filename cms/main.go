package main

import (
	"log"

	"github.com/MikelSot/metal-bat/infrastructure/handler/response"
)

func main() {
	config := newConfiguration("./configuration.json")
	api := newHTTP(config, response.HTTPErrorHandler)
	logger := newLogrus(config.LogFolder, false)

	db, err := newDBConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}
}
