package login

import (
	"bankapp/db"
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type Service interface {
	login(ctx context.Context, req userLoginRequest) (err error)
}

type loginService struct {
	store  db.LoginStorer
	logger *zap.SugaredLogger
}

func (ls *loginService) login(ctx context.Context, ul userLoginRequest) (err error) {
	fmt.Println("ul is --> ", ul.Email)
	user, err := ls.store.FindUserByEmail(ctx, ul.Email)
	// TODO: Handle the err

	//authenticate the user
	err = authenticateUser(user)
	// TODO: Handle wrong password

	// Genrate the
	err = generateJWT(user)

	fmt.Println("user is --> ", user)
	return
}

func authenticateUser(user db.Users) (err error) {
	// TODO: check if the password is correct
	return
}

func generateJWT(user db.Users) (err error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	return
}

func NewService(s db.LoginStorer, l *zap.SugaredLogger) Service {
	return &loginService{
		store:  s,
		logger: l,
	}
}
