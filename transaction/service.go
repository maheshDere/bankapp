package transaction

import (
	"bankapp/db"
	"bankapp/login"
	"bankapp/utils"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	debitAmount(ctx context.Context, req DebitCreditRequest) (balance float64, err error)
	list(ctx context.Context, d ListRequest) (response Response, err error)
	creditAmount(ctx context.Context, req DebitCreditRequest) (balance float64, err error)
}

type service struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (s *service) debitAmount(ctx context.Context, req DebitCreditRequest) (balance float64, err error) {
	err = req.Validate()
	if err != nil {
		s.logger.Warn("Error debit transaction", "err", err.Error(), "transaction", req)
		return
	}

	payload, ok := ctx.Value("claims").(*login.Claims)
	if !ok || payload.ID == "" {
		s.logger.Warn("Invalid jwt playload in debit transaction", "msg", invalidUserID.Error(), "transaction", ctx.Value("claims"))
		return balance, invalidUserID
	}

	account, err := s.store.FindAccountByUserID(ctx, payload.ID)
	if err == db.NoAccountRecordForUserID {
		s.logger.Warn("Error debit transaction", "msg", err.Error(), "transaction", req, payload)
		return balance, invalidUserID
	}

	if err != nil {
		s.logger.Error("Error debit transaction", "msg", err.Error(), "transaction", req, payload)
		return balance, err
	}

	account.UserID = payload.ID
	t := &db.Transaction{
		ID:        utils.GetUniqueId(),
		Amount:    req.Amount,
		AccountID: account.ID,
		Type:      0,
	}

	balance, err = s.store.GetTotalBalance(ctx, account.ID)
	if err != db.NoTransactions && err != nil {
		s.logger.Error("Error debit transaction", "msg", err.Error(), req, account.ID)
		return
	}

	// Checking for balance
	if (balance - req.Amount) < 0 {
		s.logger.Warn("Error debit transaction", "err", balanceLow.Error(), "transaction", req, balance)
		return balance, balanceLow
	}

	err = s.store.CreateTransaction(ctx, t)
	if err != nil {
		s.logger.Error("Error debit transaction", "msg", err.Error(), req, t)
		return
	}

	balance -= req.Amount
	return
}

func (s *service) creditAmount(ctx context.Context, req DebitCreditRequest) (balance float64, err error) {
	err = req.Validate()
	if err != nil {
		s.logger.Warn("Error credit transaction", "err", err.Error(), "transaction", req)
		return
	}

	// expecting jwt payload from ctx
	payload, ok := ctx.Value("claims").(*login.Claims)
	if !ok || payload.ID == "" {
		s.logger.Warn("Invalid jwt playload in credit transaction", "msg", invalidUserID.Error(), "transaction", ctx.Value("claims"))
		return balance, invalidUserID
	}

	account, err := s.store.FindAccountByUserID(ctx, payload.ID)
	if err == db.NoAccountRecordForUserID {
		s.logger.Error("Error credit transaction", "err", err.Error())
		return balance, invalidUserID
	}

	if err != nil {
		s.logger.Error("Error credit transaction", "msg", err.Error(), "transaction", req, payload.ID)
		return balance, err
	}

	account.UserID = payload.ID
	t := &db.Transaction{
		ID:        utils.GetUniqueId(),
		Amount:    req.Amount,
		AccountID: account.ID,
		Type:      1,
	}

	err = s.store.CreateTransaction(ctx, t)
	if err != nil {
		s.logger.Error("Error credit transaction", "msg", err.Error(), "transaction", req, t)
		return
	}

	balance, err = s.store.GetTotalBalance(ctx, account.ID)
	if err != db.NoTransactions && err != nil {
		s.logger.Error("Error credit transaction", "err", err.Error(), "transaction", req, account.ID)
		return
	}
	return
}

func (cs *service) list(ctx context.Context, req ListRequest) (response Response, err error) {
	response.Transactions = make([]db.Transaction, 0)
	fromDate, err := utils.ParseStringToTime(req.FromDate)
	if err != nil {
		cs.logger.Error("Error parsing fromDate", "err", err.Error(), "transaction", req)
		return
	}
	toDate, err := utils.ParseStringToTime(req.ToDate)
	if err != nil {
		cs.logger.Error("Error while parsing toDate", "err", err.Error(), "transcation", req)
		return
	}

	// Payload of JWT
	payload, ok := ctx.Value("claims").(*login.Claims)
	if !ok || payload.ID == "" {
		cs.logger.Warn("Invalid jwt playload in list transaction", "msg", invalidUserID.Error(), "transaction", ctx.Value("claims"))
		return
	}

	// Get account id
	account, err := cs.store.FindAccountByUserID(ctx, payload.ID)
	if err == db.NoAccountRecordForUserID {
		cs.logger.Warn("Error list transaction", "msg", err.Error(), "transaction", req, payload)
		return response, invalidUserID
	}

	if err != nil {
		cs.logger.Error("Error list transaction", "msg", err.Error(), "transaction", req, payload)
		return response, err
	}

	// Get transaction list
	transaction, err := cs.store.ListTransaction(ctx, account.ID, fromDate, toDate)
	if err == db.ErrAccountNotExist {
		cs.logger.Warn("No Account present", "err", err.Error())
		return response, errNoAccountId
	}
	if err != nil {
		cs.logger.Error("Error finding Account", "err", err.Error(), "account_id", account.ID)
		return
	}
	if transaction != nil {
		response.Transactions = transaction
	}
	return
}

func NewService(store db.Storer, logger *zap.SugaredLogger) Service {
	return &service{
		store:  store,
		logger: logger,
	}
}
