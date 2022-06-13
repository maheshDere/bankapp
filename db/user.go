package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sethvargo/go-password/password"

	"golang.org/x/crypto/bcrypt"
)

const (
	createUserQuery = `INSERT INTO users (
		id,name,email,password,role_type,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	findUserByEmailQuery = `SELECT id, name, email, password, role_type FROM users WHERE email = $1`

	createAccount = `INSERT INTO account(
		id,opening_date,user_id,created_at)
		VALUES($1,$2,$3,$4)`

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//generate random password
func generatePassword() (string, error) {
	return password.Generate(8, 2, 1, false, false)
}

//generate random id for user
func generateId() (string, error) {
	return password.Generate(10, 10, 0, false, false)
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
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByEmailQuery, email)
	})
	return
}

func (s *store) CreateUser(ctx context.Context, user *User) (err error) {
	now := time.Now()
	id, err := generateId()
	fmt.Println(id)
	accountId, _ := generateId()

	password, err := generatePassword()
	password, err = HashPassword(password)
	fmt.Println("Inside the CreateUser db  user---->", password)
	//add error handling here

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			id,
			user.Name,
			user.Email,
			password,
			user.RoleType,
			now,
			now,
		)
		if err != nil {
			return err
		}

		if user.RoleType == "customer" {
			_, err = s.db.Exec(
				createAccount,
				accountId,
				now,
				id,
				now,
			)
		}
		return err
	})
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
