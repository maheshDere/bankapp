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
	balanceQuery          = `SELECT coalesce(SUM(CASE tnx_type WHEN 1 THEN amount WHEN 0 THEN -amount), 0.00) FROM transactions WHERE account_id = $1`
	listTransactionsQuery = `SELECT id,tnx_type,amount,account_id,created_at FROM transactions WHERE account_id = $1`
)

type Transaction struct {
	ID        string    `db:"id"`
	Type      int       `db:"tnx_type"`
	Amount    float64   `db:"amount"`
	AccountID string    `db:"account_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *store) DebitTransaction(ctx context.Context, t *Transaction) (err error) {
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

func (s *store) FindTransactionsById(ctx context.Context, accountId string) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		fmt.Println(accountId)
		return s.db.SelectContext(ctx, &transactions, listTransactionsQuery, accountId)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrAccountNotExist
	}

	return
}
