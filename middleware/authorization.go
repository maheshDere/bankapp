package middleware

import (
	"bankapp/config"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accountantRoutes = []string{"createUser"}
	customerRoutes   = []string{"credit", "debit"}
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("We are here")
		token, err := readToken(r)
		claims, err := validateToken(token)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
		}
		fmt.Println("claims", claims)
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println("Came here")
		next.ServeHTTP(w, r)
	})
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

//validateToken will validate the given token, and it will return the claims or error
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
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
	jwt.StandardClaims
}
