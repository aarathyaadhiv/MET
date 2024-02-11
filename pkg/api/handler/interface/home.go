package handlerInterface

import "github.com/gin-gonic/gin"

type HomeHandler interface {
	Home(c *gin.Context)
}
