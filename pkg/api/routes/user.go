package routes

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	middleInterface "github.com/aarathyaadhiv/met/pkg/api/middleware/interface"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup, userHandler handlerInterface.UserHandler, authMiddleware middleInterface.AuthMiddleware, activityHandler handlerInterface.ActivityHandler, homeHandler handlerInterface.HomeHandler,chatHandler handlerInterface.ChatHandler) {
	route.POST("/sendOtp", userHandler.SendOtp)
	route.POST("/verify", userHandler.VerifyOtp)
	route.Use(authMiddleware.UserAuthorization())
	{
		profile := route.Group("/profile")
		{
			profile.POST("", userHandler.AddProfile)
			profile.GET("", userHandler.GetProfile)
			profile.PUT("", userHandler.UpdateUser)
		}
		like := route.Group("/like")
		{
			like.POST("/:id", activityHandler.Like)
			like.GET("", activityHandler.GetLike)
		}
		preference := route.Group("/preference")
		{
			preference.PUT("", userHandler.UpdatePreference)
			preference.GET("", userHandler.GetPreference)
		}
		route.GET("/home", homeHandler.Home)
		route.DELETE("/unmatch/:id", activityHandler.Unmatch)
		route.GET("/match", activityHandler.GetMatch)
		route.POST("/report/:id", activityHandler.Report)
		route.POST("/block/:id", activityHandler.BlockUser)
		
		chat:=route.Group("/chat")
		{
			chat.GET("",chatHandler.GetChats)
			chat.GET("/:chatId/message",chatHandler.GetMessages)
			chat.POST("/:chatId/message",chatHandler.SendMessage)
			chat.PATCH("/message/read",chatHandler.MakeMessageRead)
		}
		
	}
}
