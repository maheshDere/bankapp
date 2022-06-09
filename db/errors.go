package db

import "errors"

var (
	NoAccountRecordForUserID = errors.New("No account found for given userId")
	NoTransactions           = errors.New("No transactions found")
)
