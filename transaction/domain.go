package transaction

import "bankapp/db"

type debitRequest struct {
	Amount float64 `json:"amount"`
}

func (d debitRequest) Validate() error {
	if d.Amount <= 0 {
		return invalidAmount
	}
	return nil
}

type FindByTransactionIdResponse struct {
	Transactions []db.Transaction `json:"transactions"`
}
