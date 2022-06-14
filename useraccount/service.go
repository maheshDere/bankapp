package useraccount

import (
	"bankapp/db"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (resp db.CreateUserResponse, err error)
}

type userAccountService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (us *userAccountService) create(ctx context.Context, req createRequest) (resp db.CreateUserResponse, err error) {
	fmt.Println("Inside create user service")

	resp, err = us.store.CreateUserAccount(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		RoleType: req.RoleType,
	})

	if err != nil {
		us.logger.Error("Error while creating user", "err", err.Error())
	}

	return resp, err
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userAccountService{
		store:  s,
		logger: l,
	}
}
