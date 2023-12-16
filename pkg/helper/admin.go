package helper

import (
	"time"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/golang-jwt/jwt"
)


type AdminCustomClaim struct{
	Id uint
	Role string
	jwt.StandardClaims
}

func GenerateAdminToken(id uint)(string,error){
	claim:=&AdminCustomClaim{
		Id: id,
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute *30).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	config:=config.Config{}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenString,err:=token.SignedString([]byte(config.JwtSecret))
	return tokenString,err
}