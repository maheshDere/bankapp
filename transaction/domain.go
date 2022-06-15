package transaction

import (
	"bankapp/db"
)

type DebitCreditRequest struct {
	Amount float64 `json:"amount"`
}

func (d DebitCreditRequest) Validate() error {
	if d.Amount <= 0 {
		return invalidAmount
	}
	return nil
}

type ListRequest struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type CreateTransactionResponse struct {
	Message      string  `json:"message"`
	TotalBalance float64 `json:"balance"`
}

type Response struct {
	Transactions []db.Transaction `json:"transactions"`
}
