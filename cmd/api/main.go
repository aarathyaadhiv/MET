package main

import (
	"log"

	"github.com/aarathyaadhiv/met/cmd/api/docs"
	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/di"
	"github.com/joho/godotenv"
)

func main() {
	docs.SwaggerInfo.Title = "MET"
	docs.SwaggerInfo.Description = "Dating App"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	err := godotenv.Load()

	if err != nil {
		log.Fatal("cannot load env:", err)
	}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Error in loading config file")
	}
	server,diErr:=di.InitializeAPI(config)
	if diErr!=nil{
		log.Fatal("Error in initializing API")
	}else{
		server.Start()
	}

}
