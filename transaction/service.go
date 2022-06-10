package transaction

import (
	"bankapp/db"
	"bankapp/utils"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	debitAmount(ctx context.Context, req debitRequest) (err error)
	findByID(ctx context.Context, accId string) (response FindByTransactionIdResponse, err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (s *transactionService) debitAmount(ctx context.Context, req debitRequest) (err error) {
	err = req.Validate()
	if err != nil {
		s.logger.Error("Invalid amount for debit transaction", "err", err.Error())
		return
	}

	// expecting jwt payload from ctx
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return invalidUserID
	}

	accounts, err := s.store.FindByUserID(ctx, userID)
	if err == db.NoAccountRecordForUserID {
		s.logger.Error("No account found for the userId", "err", err.Error())
		return invalidUserID
	}

	if err != nil {
		s.logger.Error("Failed getting account details", "msg", err.Error(), req, userID)
		return err
	}

	accounts.UserID = userID
	// amount will be negative for debit
	t := &db.Transaction{
		ID:        utils.GetUniqueId(),
		Amount:    0 - req.Amount,
		AccountID: accounts.ID,
		Type:      0,
	}

	balance, err := s.store.GetTotalBalance(ctx, accounts.ID)
	if err != db.NoTransactions && err != nil {
		s.logger.Error("Error while getting account balance", "msg", err.Error(), req, accounts.ID)
		return
	}
	// Checking for balance
	if (balance - req.Amount) < 0 {
		return balanceLow
	}

	err = s.store.DebitTransaction(ctx, t)
	if err != nil {
		s.logger.Error("Error in create debit transaction", "msg", err.Error(), req, *t)
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
