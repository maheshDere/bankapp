package db

import (
	"context"
	"database/sql"
	"time"
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

const (
	uadateUserQuery = "UPDATE users SET name=$1 ,password=$2,updated_at=$3 where id=$4"
)

func (s *store) UpdateUser(ctx context.Context, user *Users) (err error) {
	now := time.Now()
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			uadateUserQuery,
			user.Name,
			user.Password,
			now,
			user.ID,
		)
		return err
	})
}
