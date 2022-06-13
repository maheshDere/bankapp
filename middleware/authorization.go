package middleware

import (
	"bankapp/api"
	"bankapp/config"
	"bankapp/login"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routName := r.URL.RequestURI()
		if routName == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		var isAccountant bool
		var isCustomer bool

		token, err := readToken(r)
		claims, err := validateToken(token)
		if err != nil {
			api.Error(w, http.StatusUnauthorized, api.Response{Message: err.Error()})
			return
		}
		if claims.RoleType == "accountant" && strings.Contains(routName, "user") {
			isAccountant = true
		}
		if claims.RoleType == "customer" && strings.Contains(routName, "transaction") {
			isCustomer = true
		}
		if isAccountant || isCustomer {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "claims", claims)))
		} else {
			api.Error(w, http.StatusForbidden, api.Response{Message: "Access Denied"})
		}
		return
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
func validateToken(jwtToken string) (login.Claims, error) {
	if jwtToken == "" {
		err := errors.New("Authorization token is missing")
		return login.Claims{}, err
	}
	// Parse the token
	token, err := jwt.ParseWithClaims(jwtToken, &login.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.InitJWTConfiguration().JwtSignature), nil
	})

	if err != nil {
		err = errors.New("Token is invalid")
		return login.Claims{}, err
	}
	claims := token.Claims.(*login.Claims)
	return *claims, err
}

/* type Claims struct {
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
	jwt.StandardClaims
}
*/
