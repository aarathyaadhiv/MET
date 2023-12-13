package main

import (
	"fmt"
	"log"

	"github.com/aarathyaadhiv/met/pkg/config"
)



func main(){
	config,configErr:=config.LoadConfig()
	if configErr!=nil{
		log.Fatal("Error in loading config file")
	}
	fmt.Println(config)
}