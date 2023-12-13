package routes

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.RouterGroup, userHandler handlerInterface.UserHandler) {
	route.POST("/sendOtp",userHandler.SendOtp)
	route.POST("/verify",userHandler.VerifyOtp)
}
