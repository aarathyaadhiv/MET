package routes

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/met/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)


func AdminRoutes(route *gin.RouterGroup,adminHandler handlerInterface.AdminHandler){
	route.POST("/signUp",adminHandler.SignUp)
	route.POST("/login",adminHandler.Login)
	route.Use(middleware.AdminAuthorization)
	{
		user:=route.Group("/users")
		{
			user.GET("",adminHandler.GetUsers)
			user.PATCH("/:id",adminHandler.BlockOrUnBlock)
		}
	}
}