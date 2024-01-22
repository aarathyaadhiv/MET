package server

import (
	"net/http"

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

func NewServerHTTP(userHandler handlerInterface.UserHandler, adminHandler handlerInterface.AdminHandler, authMiddleware middleInterface.AuthMiddleware, activityHandler handlerInterface.ActivityHandler, homeHandler handlerInterface.HomeHandler, chatHandler handlerInterface.ChatHandler,subscriptionHandler handlerInterface.SubscriptionHandler) *ServerHTTP {
	server := gin.New()
	server.LoadHTMLGlob("templates/*")
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // Render the "chat.html" template
	})
	server.Use(gin.Logger())

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(server.Group("/"), userHandler, authMiddleware, activityHandler, homeHandler, chatHandler,subscriptionHandler)
	routes.AdminRoutes(server.Group("/admin"), adminHandler, authMiddleware,subscriptionHandler)

	return &ServerHTTP{engine: server}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3001")
}
