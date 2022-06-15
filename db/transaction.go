package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	debitQuery = `INSERT INTO transactions(id, tnx_type, amount, account_id, created_at) VALUES($1, $2, $3, $4, $5)`
	// 1 is for credit and 0 for debit
	balanceQuery          = `SELECT coalesce(SUM(CASE tnx_type WHEN 1 THEN amount WHEN 0 THEN -amount ELSE 0 END), 0.00) FROM transactions WHERE account_id = $1`
	listTransactionsQuery = `SELECT id,tnx_type,amount,account_id,created_at FROM transactions WHERE account_id = $1 AND created_at::date BETWEEN $2 AND $3`
)

type Transaction struct {
	ID        string    `db:"id" json:"id"`
	Type      int       `db:"tnx_type" json:"tnx_type"`
	Amount    float64   `db:"amount" json:"amount"`
	AccountID string    `db:"account_id" json:"account_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (s *store) CreateTransaction(ctx context.Context, t *Transaction) (err error) {
	now := time.Now()
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err := s.db.Exec(
			debitQuery,
			t.ID,
			t.Type,
			t.Amount,
			t.AccountID,
			now,
		)
		return err
	})
}

func (s *store) GetTotalBalance(ctx context.Context, accountId string) (balance float64, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &balance, balanceQuery, accountId)
	})
	if err == sql.ErrNoRows {
		return balance, NoTransactions
	}

	return
}

func (s *store) ListTransaction(ctx context.Context, accountId string, fromDate, toDate time.Time) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		fmt.Println(accountId)
		return s.db.SelectContext(ctx, &transactions, listTransactionsQuery, accountId, fromDate, toDate)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrAccountNotExist
	}

	return
}
