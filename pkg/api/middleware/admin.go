package middleware

import (
	"fmt"
	"net/http"

	middleInterface "github.com/aarathyaadhiv/met/pkg/api/middleware/interface"
	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/helper"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	UserRepository interfaces.UserRepository
}

func NewAuthMiddleware(userRepo interfaces.UserRepository) middleInterface.AuthMiddleware {
	return &AuthMiddleware{
		UserRepository: userRepo,
	}
}

func (a *AuthMiddleware) AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		accessTokens, err := c.Cookie("accessAdminToken")
		if err != nil {
			errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errRes)
			c.Abort()
			return
		}
		accessToken, err := ValidateAdminToken(accessTokens)
		if err != nil || !accessToken.Valid {

			refreshTokens, err := c.Cookie("refreshAdminToken")
			if err != nil {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			refreshToken, err := ValidateAdminToken(refreshTokens)
			if err != nil || !refreshToken.Valid {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			claim, ok := accessToken.Claims.(*helper.AdminCustomClaim)
			if !ok {
				errRes := response.MakeResponse(http.StatusUnauthorized, "claim recovery failed", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			if claim.Role == "user" {
				errRes := response.MakeResponse(http.StatusUnauthorized, "it is not admin token", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			id := claim.Id

			block, err := a.UserRepository.IsBlocked(id)
			if block {
				errRes := response.MakeResponse(http.StatusUnauthorized, "claim recovery failed", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			access, refresh, err := helper.GenerateAdminToken(id)
			if err != nil {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			c.SetCookie("accessAdminToken", access, 4500, "", "", false, true)
			c.SetCookie("refreshAdminToken", refresh, 4500, "", "", false, true)
			c.Set("adminId", id)
			c.Next()
			return
		}
		claim, ok := accessToken.Claims.(*helper.AdminCustomClaim)
		if !ok {
			errRes := response.MakeResponse(http.StatusUnauthorized, "claim recovery failed", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errRes)
			c.Abort()
			return
		}
		if claim.Role == "user" {
			errRes := response.MakeResponse(http.StatusUnauthorized, "it is not admin token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errRes)
			c.Abort()
			return
		}
		id := claim.Id
		c.Set("adminId", id)
		c.Next()

	}
}

func ValidateAdminToken(tokenString string) (*jwt.Token, error) {
	config := config.Config{}
	token, err := jwt.ParseWithClaims(tokenString, &helper.AdminCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", t.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})
	return token, err
}

// func AdminAuthorization(c *gin.Context){
// 	tokenString := c.Request.Header.Get("Authorization")

// 		config := config.Config{}

// 		token, err := jwt.ParseWithClaims(tokenString, &helper.AdminCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method:%v", t.Header["alg"])
// 			}
// 			return []byte(config.JwtSecret), nil
// 		})

// 		if err != nil || !token.Valid {

// 			errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, errRes)
// 			c.Abort()
// 			return
// 		}
// 		claim, ok := token.Claims.(*helper.AdminCustomClaim)
// 		if !ok {
// 			errRes := response.MakeResponse(http.StatusUnauthorized, "claim recovery failed", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, errRes)
// 			c.Abort()
// 			return
// 		}
// 		if claim.Role == "user" {
// 			errRes := response.MakeResponse(http.StatusUnauthorized, "it is not admin token", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, errRes)
// 			c.Abort()
// 			return
// 		}
// 		id := claim.Id
// 		c.Set("adminId", id)
// 		c.Next()
// }