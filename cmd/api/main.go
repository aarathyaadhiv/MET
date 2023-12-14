package main

import (
	"log"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/di"
)

func main() {
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
