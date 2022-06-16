package login

import "github.com/golang-jwt/jwt/v4"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() (err error) {
	if l.Password == "" {
		return errEmptyPassword
	}
	if l.Email == "" {
		return errEmptyEmail
	}
	return
}

type Claims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
	jwt.StandardClaims
}
