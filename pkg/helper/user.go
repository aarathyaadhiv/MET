package helper

import (
	"time"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/golang-jwt/jwt"
)

type CustomUserClaim struct {
	Id   uint
	Role string
	jwt.StandardClaims
}

func GenerateUserToken(userId uint) (string, string, error) {
	claim := &CustomUserClaim{
		Id:   userId,
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 400).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	config := config.Config{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", "", err
	}
	claims := &CustomUserClaim{
		Id:   userId,
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 500).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	RToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := RToken.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshToken, nil
}
