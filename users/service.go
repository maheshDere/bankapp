package users

import (
	"bankapp/db"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Service interface {
	update(ctx context.Context, req updateRequest, userId string) (err error)
	create(ctx context.Context, req createRequest) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *userService) update(ctx context.Context, c updateRequest, userID string) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for category update", "err", err.Error(), "users", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.Users{
		Password: c.Password, //pass encrypt
		Name:     c.Name,
		ID:       userID,
	})
	if err != nil {
		cs.logger.Error("Error updating User", "err", err.Error(), "users", c)
		return
	}

	return
}

func (us *userService) create(ctx context.Context, req createRequest) (err error) {
	fmt.Println("Inside create user service")

	err = us.store.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		RoleType: req.RoleType,
	})

	fmt.Println(err)
	//add error handling

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
