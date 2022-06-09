package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorerMock struct {
	mock.Mock
}

func (m *StorerMock) GetTotalBalance(ctx context.Context, accountId string) (balance float64, err error) {
	args := m.Called(ctx, accountId)
	balance, _ = args.Get(0).(float64)
	return balance, args.Error(1)
}

func (m *StorerMock) DebitTransaction(ctx context.Context, t *Transaction) (err error) {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *StorerMock) FindByUserID(ctx context.Context, userID string) (acc Account, err error) {
	args := m.Called(ctx, userID)
	acc, _ = args.Get(0).(Account)
	return acc, args.Error(1)
}
