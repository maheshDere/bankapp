package db

import "errors"

var (
	ErrAccountNotExist = errors.New("Account Id does not exist in db")
)
