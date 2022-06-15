// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	db "bankapp/db"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Storer) CreateUser(ctx context.Context, user *db.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *db.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DebitTransaction provides a mock function with given fields: ctx, t
func (_m *Storer) DebitTransaction(ctx context.Context, t *db.Transaction) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *db.Transaction) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserByID provides a mock function with given fields: ctx, id
func (_m *Storer) DeleteUserByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByUserID provides a mock function with given fields: ctx, userID
func (_m *Storer) FindByUserID(ctx context.Context, userID string) (db.Account, error) {
	ret := _m.Called(ctx, userID)

	var r0 db.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Account); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransactionsById provides a mock function with given fields: ctx, accountId
func (_m *Storer) FindTransactionsById(ctx context.Context, accountId string) ([]db.Transaction, error) {
	ret := _m.Called(ctx, accountId)

	var r0 []db.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string) []db.Transaction); ok {
		r0 = rf(ctx, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByEmail provides a mock function with given fields: ctx, email
func (_m *Storer) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
	ret := _m.Called(ctx, email)

	var r0 db.User
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalBalance provides a mock function with given fields: ctx, accountId
func (_m *Storer) GetTotalBalance(ctx context.Context, accountId string) (float64, error) {
	ret := _m.Called(ctx, accountId)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string) float64); ok {
		r0 = rf(ctx, accountId)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, category
func (_m *Storer) UpdateUser(ctx context.Context, category *db.User) error {
	ret := _m.Called(ctx, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *db.User) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
