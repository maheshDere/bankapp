package user

import (
	"bankapp/db"
	"context"
	"fmt"

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
	fmt.Println("Inside create user service")
	err = us.store.CreateUser(ctx, &db.User{
		Id:        "1",
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role_type: req.RoleType,
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
