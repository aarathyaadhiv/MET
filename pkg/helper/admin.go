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

func GenerateAdminToken(id uint)(string,string,error){
	claim:=&AdminCustomClaim{
		Id: id,
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute *3).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	config:=config.Config{}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenString,err:=token.SignedString([]byte(config.JwtSecret))
	if err!=nil{
		return "","",err
	}
	claims:=&AdminCustomClaim{
		Id: id,
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour *24).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	refreshToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	refreshTokenS,err:=refreshToken.SignedString([]byte(config.JwtSecret))
	if err!=nil{
		return "","",err
	}
	return tokenString,refreshTokenS,nil
}