package db

import (
	"context"
	"database/sql"
	"time"
)

const (
	listTransactionsQuery = `SELECT 
	id,tnx_type,amount,account_id,created_at
	FROM transactions WHERE account_id = $1`
)

type Transaction struct {
	ID        string    `db:"id"`
	TnxType   string    `db:"tnx_type"`
	Amount    float64   `db:"amount"`
	AccountId string    `db:"account_id`
	CreatedAt time.Time `db:"created_at"`
}

func (s *store) FindTransactionsById(ctx context.Context, accountId string) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &transactions, listTransactionsQuery, accountId)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrAccountNotExist
	}
	return
}
