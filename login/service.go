package login

import (
	"bankapp/config"
	"bankapp/db"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var secretKey = []byte(config.JWTSignature())

type Service interface {
	login(ctx context.Context, req loginRequest) (tokenString string, err error)
}

type loginService struct {
	store  db.LoginStorer
	logger *zap.SugaredLogger
}

type Claims struct {
	Email string `json:"email"`
	Role  int    `json:"role"`
	jwt.StandardClaims
}

func (ls *loginService) login(ctx context.Context, ul loginRequest) (tokenString string, err error) {
	fmt.Println("ul is --> ", ul.Email)
	user, err := ls.store.FindUserByEmail(ctx, ul.Email)
	// TODO: Handle the err

	//authenticate the user
	err = authenticateUser(user)
	// TODO: Handle wrong password

	// Genrate the
	tokenString, err = generateJWT(user.Email, user.RoleType)
	fmt.Println(" --> ", tokenString)
	//fmt.Println("user is --> ", user)
	return
}

func authenticateUser(user db.Users) (err error) {
	// TODO: check if the password is correct
	return
}

func generateJWT(email string, roleType int) (tokenString string, err error) {
	fmt.Println("inside generateJWT email is --> ", email)
	expirationTime := time.Now().Add(time.Minute * 30).Unix()
	claims := &Claims{
		Email: email,
		Role:  roleType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
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
