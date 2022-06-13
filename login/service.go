package login

import (
	"bankapp/config"
	"bankapp/db"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(config.InitJWTConfiguration().JwtSignature)

type Service interface {
	login(ctx context.Context, req loginRequest) (tokenString string, err error)
}

type loginService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

type Claims struct {
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
	jwt.StandardClaims
}

func (ls *loginService) login(ctx context.Context, ul loginRequest) (tokenString string, err error) {
	user, err := ls.store.FindUserByEmail(ctx, ul.Email)
	// TODO: Handle the err
	if user.Email == "" {
		err = errors.New("Invalid Email or Password")
		return
	}
	// Authenticate the user
	matched := authenticateUser(user, ul.Password)
	// TODO: Handle wrong password
	if !matched {
		err = errors.New("Invalid Email or Password")
		return
	}
	// Generate the
	tokenString, err = generateJWT(user.Email, user.RoleType)
	return
}

func authenticateUser(user db.User, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

func generateJWT(email, roleType string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.InitJWTConfiguration().TokenExpiry)).Unix()
	claims := &Claims{
		Email:    email,
		RoleType: roleType,
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

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &loginService{
		store:  s,
		logger: l,
	}
}
