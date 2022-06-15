package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	findUserByEmailQuery = `SELECT * FROM users WHERE email = $1`

	deleteUserByIDQuery = `DELETE FROM users WHERE id = $1`
	updateUserQuery     = "UPDATE users SET name=$1 ,password=$2,updated_at=$3 where id=$4"
)

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	RoleType string `db:"role_type"`
}

func (s *store) UpdateUser(ctx context.Context, user *User) (err error) {
	now := time.Now()
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			user.Name,
			user.Password,
			now,
			user.ID,
		)
		return err
	})
}

func (s *store) FindUserByEmail(ctx context.Context, email string) (user User, err error) {
	fmt.Println("Inside FindUserByEmail method db")
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByEmailQuery, email)
	})
	// TODO: handle error
	return
}

func (s *store) DeleteUserByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteUserByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrUserNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}
