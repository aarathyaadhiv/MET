package routes

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	middleInterface "github.com/aarathyaadhiv/met/pkg/api/middleware/interface"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup, userHandler handlerInterface.UserHandler, authMiddleware middleInterface.AuthMiddleware) {
	route.POST("/sendOtp", userHandler.SendOtp)
	route.POST("/verify", userHandler.VerifyOtp)
	route.Use(authMiddleware.UserAuthorization())
	{
		profile := route.Group("/profile")
		{
			profile.POST("", userHandler.AddProfile)
			profile.GET("", userHandler.GetProfile)
		}
	}
}
