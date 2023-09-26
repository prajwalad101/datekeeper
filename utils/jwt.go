package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateAccessToken(userId int, iat time.Time) (string, error) {
	claims := &JwtCustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(iat.Add(time.Hour * 24 * 7)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(Env.JWTSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Authorize(requestToken string, secret string) (*JwtCustomClaims, error) {
	if len(strings.Split(requestToken, ".")) != 3 {
		return nil, errors.New("Invalid Token Construct")
	}

	token, err := jwt.ParseWithClaims(
		requestToken,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)

	if !ok {
		return nil, fmt.Errorf("Permission denied")
	}

	if !token.Valid {
		return nil, fmt.Errorf("Permission denied")
	}
	return claims, nil
}
