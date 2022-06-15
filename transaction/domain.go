package transaction

import (
	"bankapp/db"
)

type debitCreditRequest struct {
	Amount float64 `json:"amount"`
}

func (d debitCreditRequest) Validate() error {
	if d.Amount <= 0 {
		return invalidAmount
	}
	return nil
}

type listRequest struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type createTransactionResponse struct {
	Message      string  `json:"message"`
	TotalBalance float64 `json:"balance"`
}

type Response struct {
	Transactions []db.Transaction `json:"transactions"`
}
