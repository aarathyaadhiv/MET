package middleware

import (
	"fmt"
	"net/http"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/helper"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// func UserAuthorization(c *gin.Context) {
// 	tokenString := c.Request.Header.Get("Authorization")
// 	config := config.Config{}

//		token, err := jwt.ParseWithClaims(tokenString, &helper.CustomUserClaim{}, func(t *jwt.Token) (interface{}, error) {
//			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected error function: %v", t.Header["alg"])
//			}
//			return []byte(config.JwtSecret), nil
//		})
//		if err != nil || !token.Valid {
//			errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
//			c.JSON(http.StatusUnauthorized, errRes)
//			c.Abort()
//			return
//		}
//		claim, ok := token.Claims.(*helper.CustomUserClaim)
//		if !ok {
//			errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
//			c.JSON(http.StatusUnauthorized, errRes)
//			c.Abort()
//			return
//		}
//		id := claim.Id
//		c.Set("userId", id)
//		c.Next()
//	}
func (a *AuthMiddleware) UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessTokens, err := c.Cookie("accessToken")

		if err != nil {
			errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errRes)
			c.Abort()
			return
		}
		accessToken, err := ValidateUserToken(accessTokens)
		if err != nil || !accessToken.Valid {

			refreshTokens, err := c.Cookie("refreshToken")
			if err != nil {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			refreshToken, err := ValidateUserToken(refreshTokens)
			if err != nil || !refreshToken.Valid {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			claim, ok := refreshToken.Claims.(*helper.CustomUserClaim)
			if !ok {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, "it is not admin token")
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			id := claim.Id

			block, err := a.UserRepository.IsBlocked(id)
			if err != nil {
				errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, "error in block checking from repository")
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			if block {

				errRes := response.MakeResponse(http.StatusUnauthorized, "claim recovery failed", nil, "user is blocked")
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			access, refresh, err := helper.GenerateUserToken(id)
			if err != nil {
				errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
				c.JSON(http.StatusUnauthorized, errRes)
				c.Abort()
				return
			}

			c.SetCookie("accessToken", access, 172800, "", "", false, true)
			c.SetCookie("refreshToken", refresh, 172800, "", "", false, true)
			c.Set("userId", id)
			c.Next()
			return
		}
		claim, ok := accessToken.Claims.(*helper.CustomUserClaim)
		if !ok {
			errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, "error in claim recovery")
			c.JSON(http.StatusUnauthorized, errRes)
			c.Abort()
			return
		}
		id := claim.Id
		c.Set("userId", id)
		c.Next()

	}
}

func ValidateUserToken(tokenString string) (*jwt.Token, error) {
	config := config.Config{}
	token, err := jwt.ParseWithClaims(tokenString, &helper.CustomUserClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", t.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})
	return token, err
}
