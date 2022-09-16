package main

import (
	"fmt"
	"log"

	"github.com/MikelSot/metal-bat/infrastructure/handler"
	"github.com/MikelSot/metal-bat/infrastructure/handler/response"
	"github.com/MikelSot/metal-bat/model"
)

func main() {
	config := newConfiguration("./configuration.json")
	api := newHTTP(config, response.HTTPErrorHandler)
	logger := newLogrus(config.LogFolder, false)
	loadSignatures(config, logger)

	db, err := newDBConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(model.RouterSpecification{
		Config:   config,
		Api:      api,
		Logger:   logger,
		DBEngine: config.Database.Engine,
		DB:       db,
	})

	port := fmt.Sprintf(":%d", config.ServerPort)
	if config.IsHttps {
		log.Fatal(api.StartTLS(port, config.CertPem, config.KeyPem))
		return
	}

	log.Println("Starting service")
	log.Fatal(api.Start(port))
}
