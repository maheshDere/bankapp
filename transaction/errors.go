package transaction

import "errors"

var (
	invalidAmount  = errors.New("Please enter a valid amount")
	invalidUserID  = errors.New("No user found")
	balanceLow     = errors.New("You don't have enough balance to proceed")
	errNoAccountId = errors.New(" Account is not present")
)
