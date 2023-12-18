package server

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/met/pkg/api/routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)


type ServerHTTP struct{
	engine *gin.Engine
}

func NewServerHTTP(userHandler handlerInterface.UserHandler,adminHandler handlerInterface.AdminHandler)*ServerHTTP{
	server:=gin.New()
	server.Use(gin.Logger())

	server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(server.Group("/"),userHandler)
	routes.AdminRoutes(server.Group("/admin"),adminHandler)
	
	return &ServerHTTP{engine: server}
}

func (s *ServerHTTP) Start(){
	s.engine.Run(":3001")
}