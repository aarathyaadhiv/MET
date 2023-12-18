package routes

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/met/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup, userHandler handlerInterface.UserHandler) {
	route.POST("/sendOtp",userHandler.SendOtp)
	route.POST("/verify",userHandler.VerifyOtp)
	route.Use(middleware.UserAuthorization)
	{
		profile:=route.Group("/profile")
		{
			profile.POST("",userHandler.AddProfile)
			profile.GET("",userHandler.GetProfile)
		}
	}
}
