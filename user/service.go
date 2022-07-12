package user

import (
	"bankapp/db"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	update(ctx context.Context, req updateRequest, userId string) (err error)
	deleteByID(ctx context.Context, id string) (err error)
	//rak
	listAllUsers(ctx context.Context) (users []db.User, err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *userService) update(ctx context.Context, c updateRequest, userID string) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "users", c)
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

//rak
//list users service
func (cs *userService) listAllUsers(ctx context.Context) (users []db.User, err error) {
	// var users []db.User
	// dbUsers, err := cs.store.ListUsers(ctx)

	// for _, u := range dbUsers {
	// 	if u.Email == "accountant@bank.com" {
	// 		continue
	// 	}
	// 	u.Password = "reducted"
	// 	dbUsers = append(dbUsers, u)
	// }
	users, err = cs.store.ListUsers(ctx)

	if err == db.ErrNoUserExist {
		cs.logger.Error("User not present in db", "err", err.Error(), "users", users)
		return
	}
	if err != nil {
		cs.logger.Error("Error fetching data", "err", err.Error(), "users", users)
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
