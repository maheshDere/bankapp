package user

import "errors"

var (
	errEmptyPassword = errors.New("User password must be presnet")
	errEmptyName     = errors.New("User Name  must be presnet")
	errNoUserId      = errors.New("User is not present")
	invalidUserID    = errors.New("No user found")
	errNoAccountId   = errors.New("User with ID not exist")
)
