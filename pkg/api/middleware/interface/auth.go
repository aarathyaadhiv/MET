package middleInterface

import "github.com/gin-gonic/gin"


type AuthMiddleware interface{
	AdminAuthorization() gin.HandlerFunc
	UserAuthorization() gin.HandlerFunc
}