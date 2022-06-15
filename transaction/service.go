package transaction

import (
	"bankapp/db"
	"bankapp/utils"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Service interface {
	debitAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error)
	list(ctx context.Context, accId string, d listRequest) (response Response, err error)
	creditAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error)
}

type service struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (s *service) debitAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error) {
	err = req.Validate()
	if err != nil {
		s.logger.Error("Invalid amount for debit transaction", "err", err.Error())
		return
	}

	// expecting jwt payload from ctx
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		s.logger.Error("Invalid user Id in jwt payload", "err", invalidUserID.Error())
		return balance, invalidUserID
	}

	account, err := s.store.FindAccountByUserID(ctx, userID)
	if err == db.NoAccountRecordForUserID {
		s.logger.Error("No account found for the userId", "err", err.Error())
		return balance, invalidUserID
	}

	if err != nil {
		s.logger.Error("Failed getting account details", "msg", err.Error(), req, userID)
		return balance, err
	}

	account.UserID = userID
	t := &db.Transaction{
		ID:        utils.GetUniqueId(),
		Amount:    req.Amount,
		AccountID: account.ID,
		Type:      0,
	}

	balance, err = s.store.GetTotalBalance(ctx, account.ID)
	if err != db.NoTransactions && err != nil {
		s.logger.Error("Error while getting account balance", "msg", err.Error(), req, account.ID)
		return
	}
	// Checking for balance
	if (balance - req.Amount) < 0 {
		s.logger.Error("Insufficient funds for debit", "err", balanceLow.Error())
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

func (s *service) creditAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error) {
	err = req.Validate()
	if err != nil {
		s.logger.Error("Invalid amount for debit transaction", "err", err.Error())
		return
	}

	// expecting jwt payload from ctx
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		s.logger.Error("Invalid user Id in jwt payload", "err", invalidUserID.Error())
		return balance, invalidUserID
	}

	account, err := s.store.FindAccountByUserID(ctx, userID)
	if err == db.NoAccountRecordForUserID {
		s.logger.Error("No account found for the userId", "err", err.Error())
		return balance, invalidUserID
	}

	if err != nil {
		s.logger.Error("Failed getting account details", "msg", err.Error(), req, userID)
		return balance, err
	}

	account.UserID = userID
	t := &db.Transaction{
		ID:        utils.GetUniqueId(),
		Amount:    req.Amount,
		AccountID: account.ID,
		Type:      1,
	}

	err = s.store.CreateTransaction(ctx, t)
	if err != nil {
		s.logger.Error("Error credit transaction", "msg", err.Error(), req, t)
		return
	}

	balance, err = s.store.GetTotalBalance(ctx, account.ID)
	if err != db.NoTransactions && err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting account balance in credit transaction", "msg", err.Error(), req, account.ID)
		return
	}
	return
}

func (cs *service) list(ctx context.Context, accountId string, req listRequest) (response Response, err error) {
	response.Transactions = make([]db.Transaction, 0)
	fromDate, err := utils.ParseStringToTime(req.FromDate)
	if err != nil {
		cs.logger.Error("Error while parsing", "err", err.Error())
		return
	}
	toDate, err := utils.ParseStringToTime(req.ToDate)
	if err != nil {
		cs.logger.Error("Error while parsing", "err", err.Error())
		return
	}

	transaction, err := cs.store.ListTransaction(ctx, accountId, fromDate, toDate)
	if err == db.ErrAccountNotExist {
		cs.logger.Warn("No Account present", "err", err.Error())
		return response, errNoAccountId
	}
	if err != nil {
		cs.logger.Error("Error finding Account", "err", err.Error(), "account_id", accountId)
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
