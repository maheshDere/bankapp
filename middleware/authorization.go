package middleware

import (
	"bankapp/config"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	accountantRoutes = []string{"createUser"}
	customerRoutes   = []string{"credit", "debit"}
)

func AuthorizationMiddleware(next http.HandlerFunc, routName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountantRouteFound string
		var customerRouteFound string

		token, err := readToken(r)
		claims, err := validateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if claims.RoleType == "accountant" {
			for _, accountantRouteFound = range accountantRoutes {
				if routName == accountantRouteFound {
					break
				}
			}
		}
		if claims.RoleType == "customer" {
			for _, customerRouteFound = range customerRoutes {
				if routName == customerRouteFound {
					break
				}
			}
		}

		switch {
		case accountantRouteFound != "":
			next.ServeHTTP(w, r)
		case customerRouteFound != "":
			next.ServeHTTP(w, r)
		}
		w.WriteHeader(http.StatusUnauthorized)
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
