package db

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	createUserQuery = `INSERT INTO users (
		id,name,email,password,role_type)
		VALUES($1,$2,$3,$4,$5)
	`
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role_type string `json:"role_type"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *store) CreateUser(ctx context.Context, user *User) (err error) {

	password, err := HashPassword(user.Password)
	fmt.Println("Inside the CreateUser db  user---->", password)
	//add error handling here

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.Id, user.Name, user.Email, password, user.Role_type,
		)
		return err
	})
}
