package authorization

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"service-manager/config"
	"time"
)

type Token struct {
}

func (token Token) Generate(id string) (string, error) {
	conf := config.Load()
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        id,
	})
	signedString, err := t.SignedString([]byte(conf.Http.Token))
	if err != nil {
		return "", err
	}
	return signedString, err
}

func (token Token) Analyze(t string) (*jwt.RegisteredClaims, error) {
	var claims jwt.RegisteredClaims
	withClaims, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		conf := config.Load()
		return []byte(conf.Http.Token), nil
	})
	if err != nil {
		return nil, err
	}
	if !withClaims.Valid {
		return nil, fmt.Errorf("token not valided")
	}
	registeredClaims, ok := withClaims.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("token not valided")
	}
	return registeredClaims, nil
}

func NewToken() Token {
	return Token{}
}
