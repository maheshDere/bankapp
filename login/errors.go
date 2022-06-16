package login

import "errors"

var (
	errEmptyPassword = errors.New("Password must be present")
	errEmptyEmail    = errors.New("Email must be present")
)
