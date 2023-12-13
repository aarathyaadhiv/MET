package handlerInterface

import "github.com/gin-gonic/gin"





type UserHandler interface{
	SendOtp(c *gin.Context)
	VerifyOtp(c *gin.Context)
}