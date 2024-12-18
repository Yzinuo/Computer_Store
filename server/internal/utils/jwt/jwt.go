package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Myclaims struct {
	UserId int 		`json:"user_id"`
	jwt.RegisteredClaims
}

func GenToken(secret, issuer string, expireHour, userId int) (string, error) {
	claims := Myclaims{
		UserId:userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: issuer,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}