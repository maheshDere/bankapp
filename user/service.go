package user

import (
	"bankapp/db"
	"bankapp/login"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	update(ctx context.Context, req updateRequest, userId string) (err error)
	deleteByID(ctx context.Context, id string) (err error)
	//rak
	listAllUsers(ctx context.Context) (users []db.User, err error)
	getUserByID(ctx context.Context, id string) (user db.User, err error)
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
	users = make([]db.User, 0)
	dbUsers, err := cs.store.ListUsers(ctx)

	for _, u := range dbUsers {
		if u.Email == "accountant@bank.com" {
			continue
		}
		u.Password = "reducted"
		users = append(users, u)
	}
	// users, err = cs.store.ListUsers(ctx)

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

//get user by id:
// func (cs *userService) getUserById(ctx context.Context, id string) (user db.User, err error) {

// 	// user = db.User{}
// 	dbUser, err := cs.store.GetUser(ctx, id)

// 	dbUser.Password = "Reducted"
// 	user = dbUser

// 	if err == db.ErrUserNotExist {
// 		cs.logger.Error("User not present in db", "err", err.Error(), "user", user)
// 		return
// 	}
// 	if err != nil {
// 		cs.logger.Error("Error fetching data", "err", err.Error(), "user", user)
// 		return
// 	}
// 	return
// }

func (cs *userService) getUserByID(ctx context.Context, id string) (user db.User, err error) {

	// payload of jwt
	payload, ok := ctx.Value("claims").(*login.Claims)
	if !ok || payload.ID == "" {
		cs.logger.Warn("Invalid jwt playload in get user", "msg", invalidUserID.Error(), "user", ctx.Value("claims"))
		return
	}
	//get accountant id
	user, err = cs.store.GetUser(ctx, payload.ID)
	if err == db.NoAccountRecordForUserID {
		cs.logger.Warn("Error get user", "msg", err.Error(), "user", user, payload)
		return
	}
	if err != nil {
		cs.logger.Error("Error get user", "msg", err.Error(), "user", user, payload)
		return
	}
	user.Password = "reducted"
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
