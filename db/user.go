package db

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

const (
	createUserQuery = `INSERT INTO user (
		id,name,email,password,role_type)
		VALUES($1,$2,$3,$4,$5)
	`
)

type user struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role_type int    `json:"role_type"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *store) CreateUser(ctx context.Context, user *user) (err error) {

	password, err := HashPassword(user.Password)

	//add error handling here

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.Id, user.Name, user.Email, password, user.Role_type,
		)
		return err
	})
}