package middleware

import (
	"bankapp/config"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthorizationMiddleware(next http.HandlerFunc, forAccountant bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := readToken(r)
		// TODO: handle the err

		claims, err := validateToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if forAccountant && claims.IsAccountant == true {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}
}

//readToken method will read the Authorization header and will return the token string or error
func readToken(r *http.Request) (token string, err error) {
	//TODO: header missing error
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		err = errors.New("Malform header")
		return
	}
	token = authHeader[1]
	return
}

func validateToken(jwtToken string) (Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.InitJWTConfiguration().JwtSignature), nil
	})
	// TODO: handle the err

	claims := token.Claims.(*Claims)
	return *claims, err
}

type Claims struct {
	Email        string `json:"email"`
	IsAccountant bool   `json:"isAccountant"`
	jwt.StandardClaims
}
