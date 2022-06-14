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
	findByID(ctx context.Context, accId string) (response FindByTransactionIdResponse, err error)
	creditAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (s *transactionService) debitAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error) {
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

func (s *transactionService) creditAmount(ctx context.Context, req debitCreditRequest) (balance float64, err error) {
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

func (cs *transactionService) findByID(ctx context.Context, accountId string) (response FindByTransactionIdResponse, err error) {
	transaction, err := cs.store.FindTransactionsById(ctx, accountId)
	if err == db.ErrAccountNotExist {
		cs.logger.Error("No Account present", "err", err.Error())
		return response, errNoAccountId
	}
	if err != nil {
		cs.logger.Error("Error finding Account", "err", err.Error(), "account_id", accountId)
		return
	}

	response.Transactions = transaction
	return
}

func NewService(store db.Storer, logger *zap.SugaredLogger) Service {
	return &transactionService{
		store:  store,
		logger: logger,
	}
}
