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
	createAccount = `INSERT INTO account(
		id,opening_date,user_id,created_at)
		VALUES($1,$2,$3,$4)`
)

type User struct {
	Id       string `db:"id"`
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
			id, user.Name, user.Email, password, user.RoleType, now, now,
		)
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			createAccount,
			accountId,
			now,
			id,
			now,
		)
		return err
	})
}
