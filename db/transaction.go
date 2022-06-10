package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	listTransactionsQuery = `SELECT id,tnx_type,amount,account_id,created_at FROM transactions WHERE account_id = $1`
)

type Transaction struct {
	ID        string    `db:"id"`
	TnxType   string    `db:"tnx_type"`
	Amount    float64   `db:"amount"`
	AccountID string    `db:"account_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *store) FindTransactionsById(ctx context.Context, accountId string) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		fmt.Println(accountId)
		return s.db.SelectContext(ctx, &transactions, listTransactionsQuery, accountId)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrAccountNotExist
	}
	fmt.Println("err", err)
	return
}
