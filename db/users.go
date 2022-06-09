package db

import (
	"context"
	"fmt"
	"time"
)

const (
	findUserByEmailQuery = `SELECT * FROM users WHERE email = $1`
)

type Users struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	RoleType  int       `db:"role_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (s *store) FindUserByEmail(ctx context.Context, email string) (user Users, err error) {
	fmt.Println("Inside FindUserByEmail method db")
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByEmailQuery, email)
	})
	// TODO: handle error
	return
}
