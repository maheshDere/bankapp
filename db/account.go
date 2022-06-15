package db

import (
	"context"
	"database/sql"
)

const (
	findByUserID = `SELECT id, opening_date, created_at FROM account WHERE user_id = $1`
)

type Account struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	OpeningDate string `db:"opening_date"`
	CreatedAt   string `db:"created_at"`
}

func (s *store) FindAccountByUserID(ctx context.Context, userID string) (acc Account, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &acc, findByUserID, userID)
	})
	if err == sql.ErrNoRows {
		return acc, NoAccountRecordForUserID
	}
	return
}
