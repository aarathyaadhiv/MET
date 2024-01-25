package handlerInterface

import "github.com/gin-gonic/gin"

type ChatHandler interface {
	GetChats(c *gin.Context)
	GetMessages(c *gin.Context)
	// SendMessage(c *gin.Context)
	MakeMessageRead(c *gin.Context)
	Chat(c *gin.Context)
	ChatPage(c *gin.Context)
}
