package user

import (
	"bankapp/db"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Service interface {
	update(ctx context.Context, req updateRequest, userId string) (err error)
	create(ctx context.Context, req createRequest) (resp db.CreateUserResponse, err error)
	deleteByID(ctx context.Context, id string) (err error)
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

func (us *userService) create(ctx context.Context, req createRequest) (resp db.CreateUserResponse, err error) {
	fmt.Println("Inside create user service")

	resp, err = us.store.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		RoleType: req.RoleType,
	})

	if err != nil {
		us.logger.Error("Error while creating user", "err", err.Error())
	}

	return resp, err
}

func (cs *userService) deleteByID(ctx context.Context, id string) (err error) {
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
