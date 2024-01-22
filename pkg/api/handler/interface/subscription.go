package handlerInterface

import "github.com/gin-gonic/gin"

type SubscriptionHandler interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	ActivateOrDeactivate(c *gin.Context)
	Get(c *gin.Context)
	GetById(c *gin.Context)
	GetToUsers(c *gin.Context)
	GetByIdToUsers(c *gin.Context)
	AddOrder(c *gin.Context)
	MakePayment(c *gin.Context)
	VerifyPayment(c *gin.Context)
}
