package user

import (
	"bankapp/db"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (us *userService) create(ctx context.Context, req createRequest) (err error) {
	err = us.store.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		RoleType: req.RoleType,
	})

	//add error handling

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
