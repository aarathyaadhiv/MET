package server

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	middleInterface "github.com/aarathyaadhiv/met/pkg/api/middleware/interface"
	"github.com/aarathyaadhiv/met/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handlerInterface.UserHandler, adminHandler handlerInterface.AdminHandler, authMiddleware middleInterface.AuthMiddleware,activityHandler handlerInterface.ActivityHandler,homeHandler handlerInterface.HomeHandler,chatHandler handlerInterface.ChatHandler) *ServerHTTP {
	server := gin.New()
	server.Use(gin.Logger())

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(server.Group("/"), userHandler,authMiddleware,activityHandler,homeHandler,chatHandler)
	routes.AdminRoutes(server.Group("/admin"), adminHandler, authMiddleware)

	return &ServerHTTP{engine: server}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3001")
}
