package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

const (
	findUserByEmailQuery = `SELECT * FROM users WHERE email = $1`
	createUserQuery      = `INSERT INTO users (
		id,name,email,password,role_type,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	createAccount = `INSERT INTO account(
		id,opening_date,user_id,created_at)
		VALUES($1,$2,$3,$4)`
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	RoleType  string    `db:"role_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (s *store) FindUserByEmail(ctx context.Context, email string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByEmailQuery, email)
	})
	// TODO: handle error
	return
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
	accountId, _ := generateId()

	password, err := generatePassword()
	password, err = HashPassword(password)
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
