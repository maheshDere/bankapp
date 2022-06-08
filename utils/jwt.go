package utils

import (
	"bankapp/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtPayload struct {
	Email  string `json:"email"`
	UserID int    `json:"userId"`
	Role   int    `json:"role"`
}

type JwtClaims struct {
	JwtPayload
	jwt.StandardClaims
}

func Create(email string, userId, role int) (token string, err error) {
	// key to encrypt token
	key := config.GetJwtKey()
	tokenExpiresAt := time.Now().Add(15 * time.Minute)
	// adding customer fields
	claims := JwtClaims{
		JwtPayload: JwtPayload{
			Email:  email,
			UserID: userId,
			Role:   role,
		},
		// token will expire in
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiresAt.Unix(),
		},
	}

	val := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = val.SignedString([]byte(key))
	return
}

func Validate(token string) (payload JwtPayload, err error) {
	key := config.GetJwtKey()

	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := parsedToken.Claims.(*JwtClaims)
	if ok && parsedToken.Valid {
		return claims.JwtPayload, nil
	}
	return payload, invalidToken
}
