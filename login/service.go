package login

import (
	"bankapp/config"
	"bankapp/db"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var secretKey = []byte(config.InitJWTConfiguration().JwtSignature)

type Service interface {
	login(ctx context.Context, req loginRequest) (tokenString string, err error)
}

type loginService struct {
	store  db.LoginStorer
	logger *zap.SugaredLogger
}

type Claims struct {
	Email        string `json:"email"`
	IsAccountant bool   `json:"isAccountant"`
	jwt.StandardClaims
}

func (ls *loginService) login(ctx context.Context, ul loginRequest) (tokenString string, err error) {
	user, err := ls.store.FindUserByEmail(ctx, ul.Email)
	// TODO: Handle the err

	//authenticate the user
	err = authenticateUser(user)
	// TODO: Handle wrong password

	// Generate the
	tokenString, err = generateJWT(user.Email, user.RoleType)
	return
}

func authenticateUser(user db.User) (err error) {
	// TODO: check if the password is correct
	return
}

func generateJWT(email, roleType string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.InitJWTConfiguration().TokenExpiry)).Unix()
	isAccountant := false
	if roleType == "accountant" {
		isAccountant = true
	}
	claims := &Claims{
		Email:        email,
		IsAccountant: isAccountant,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.InitJWTConfiguration().JwtSignature))
	if err != nil {
		return "", err
	}
	return
}

func NewService(s db.LoginStorer, l *zap.SugaredLogger) Service {
	return &loginService{
		store:  s,
		logger: l,
	}
}
