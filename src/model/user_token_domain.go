package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		errorMessage := "Error generating token"
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("%s: %s", errorMessage, err.Error()),
		)
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)
	token, err := jwt.Parse(RemoveBearerPrefix(tokenString), func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	token = strings.TrimPrefix(token, "Bearer ")
	return token
}
