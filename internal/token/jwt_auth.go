package token

import (
	"errors"
	"strings"
	"time"

	"avito-shop-service/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	secret string
}

func New(secret string) *JWTToken {
	return &JWTToken{
		secret: secret,
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (t *JWTToken) CreateToken(user domain.User) (string, error) {
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}

func (t *JWTToken) ParseToken(header string) (string, error) {
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 1 {
		return "", errors.New("invalid token format")
	}

	token, err := jwt.ParseWithClaims(headerParts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("unexpected signing method")
		}
		return []byte(t.secret), nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return claims.Username, nil
}
