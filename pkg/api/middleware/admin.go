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


func AdminAuthorization(c *gin.Context){
	tokenString:=c.Request.Header.Get("Authorization")
	config:=config.Config{}

	token,err:=jwt.ParseWithClaims(tokenString,&helper.AdminCustomClaim{},func(t *jwt.Token) (interface{}, error) {
		if _,ok:=t.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("unexpected signing method:%v",t.Header["alg"])
		}
		return []byte(config.JwtSecret),nil
	})

	if err!=nil || !token.Valid{
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorized", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	claim,ok:=token.Claims.(*helper.AdminCustomClaim)
	if !ok{
		errRes := response.MakeResponse(http.StatusUnauthorized, "claim recovery failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	if claim.Role=="user"{
		errRes := response.MakeResponse(http.StatusUnauthorized, "it is not admin token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	id:=claim.Id
	c.Set("adminId",id)
	c.Next()
}