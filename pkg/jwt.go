package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const TokenExpiredDuration = time.Hour * 2

var MySecret = []byte("cisco")

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

func GenToken(username string) (string, error) {

	c := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiredDuration).Unix(),
			Issuer:    "cisco46589",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(MySecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
