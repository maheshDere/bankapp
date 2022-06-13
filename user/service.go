package user

import (
	"bankapp/db"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (resp db.CreateAccountResponse, err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (us *userService) create(ctx context.Context, req createRequest) (resp db.CreateAccountResponse, err error) {

	resp, err = us.store.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		RoleType: req.RoleType,
	})

	if err != nil {
		us.logger.Error("Error while creating customer", "err", err.Error())
		return
	}

	return resp, err
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
