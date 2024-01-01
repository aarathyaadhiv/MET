package handlerInterface

import "github.com/gin-gonic/gin"


type ActivityHandler interface{
	Like(c *gin.Context)
	GetLike(c *gin.Context)
	Unmatch(c *gin.Context)
	GetMatch(c *gin.Context)
	Report(c *gin.Context)
	BlockUser(c *gin.Context)
}