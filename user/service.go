package user

import (
	"bankapp/db"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Service interface {
	Update(ctx context.Context, req UpdateRequest, userId string) (err error)
	Create(ctx context.Context, req CreateRequest) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *userService) Update(ctx context.Context, c UpdateRequest, userID string) (err error) {
	err = errors.New("validation error")
	if err != nil {
		cs.logger.Error("Invalid Request for category update", "err", err.Error(), "users", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
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

func (us *userService) Create(ctx context.Context, req CreateRequest) (err error) {
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

func (cs *userService) DeleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("User Not present", "err", err.Error(), "user_id", id)
		return errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error deleting user", "err", err.Error(), "user_id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
