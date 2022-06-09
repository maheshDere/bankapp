package transaction

import "bankapp/db"

type FindByTransactionIdResponse struct {
	Transactions []db.Transaction `json:"transactions"`
}
