package transaction

import (
	"bankapp/db"
)

type debitRequest struct {
	Amount float64 `json:"amount"`
}

func (d debitRequest) Validate() error {
	if d.Amount <= 0 {
		return invalidAmount
	}
	return nil
}

type listRequest struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type Response struct {
	Transactions []db.Transaction `json:"transactions"`
}
