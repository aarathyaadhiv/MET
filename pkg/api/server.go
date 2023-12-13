package server

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/met/pkg/api/routes"
	"github.com/gin-gonic/gin"
)


type ServerHTTP struct{
	engine *gin.Engine
}

func NewServerHTTP(userHandler handlerInterface.UserHandler)*ServerHTTP{
	server:=gin.New()
	routes.UserRoutes(server.Group("/"),userHandler)
	
	return &ServerHTTP{engine: server}
}

func (s *ServerHTTP) Start(){
	s.engine.Run(":3000")
}