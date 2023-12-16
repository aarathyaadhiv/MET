package handlerInterface

import "github.com/gin-gonic/gin"


type AdminHandler interface{
	SignUp(c *gin.Context) 
	Login(c *gin.Context)
	BlockOrUnBlock(c *gin.Context) 
	GetUsers(c *gin.Context)
}