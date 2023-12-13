package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)


type CustomUserClaim struct{
	Id uint
	Role string
	jwt.StandardClaims
}

func GenerateUserToken(userId uint)(string,error){
	claim:=&CustomUserClaim{
		Id: userId,
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute *30).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenString,err:=token.SignedString([]byte("met"))
	return tokenString,err
}