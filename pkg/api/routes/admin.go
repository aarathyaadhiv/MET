package routes

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	middleInterface "github.com/aarathyaadhiv/met/pkg/api/middleware/interface"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(route *gin.RouterGroup, adminHandler handlerInterface.AdminHandler, authMiddleware middleInterface.AuthMiddleware) {
	route.POST("/signUp", adminHandler.SignUp)
	route.POST("/login", adminHandler.Login)
	
	route.Use(authMiddleware.AdminAuthorization())
	{
		user := route.Group("/users")
		{
			user.GET("", adminHandler.GetUsers)
			user.GET("/:id",adminHandler.GetSingleUser)
			user.PATCH("/:id", adminHandler.BlockOrUnBlock)
		}
	}
}
