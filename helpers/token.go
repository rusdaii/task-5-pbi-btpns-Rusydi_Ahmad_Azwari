package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var tokenKey = []byte(os.Getenv("TOKEN_SECRET"))

type TokenClaim struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id string, username string) (tokenString string, err error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &TokenClaim{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(tokenKey)

	return
}

func ValidateToken(signedToken string) (*TokenClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*TokenClaim)

	if !ok {

		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
