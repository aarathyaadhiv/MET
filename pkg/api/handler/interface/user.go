package handlerInterface

import "github.com/gin-gonic/gin"





type UserHandler interface{
	SendOtp(c *gin.Context)
	VerifyOtp(c *gin.Context)
	AddProfile(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdatePreference(c *gin.Context)
	GetPreference(c *gin.Context)
}