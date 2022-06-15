package transaction

import (
	"bankapp/db"
	"bankapp/utils"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	debitAmount(ctx context.Context, req debitRequest) (err error)
	list(ctx context.Context, accId string, d listRequest) (response Response, err error)
}

type service struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (s *service) debitAmount(ctx context.Context, req debitRequest) (err error) {
	err = req.Validate()
	if err != nil {
		s.logger.Error("Invalid amount for debit transaction", "err", err.Error())
		return
	}

	// expecting jwt payload from ctx
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		s.logger.Error("Invalid user Id in jwt payload", "err", invalidUserID.Error())
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
	t := &db.Transaction{
		ID:        utils.GetUniqueId(),
		Amount:    req.Amount,
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
		s.logger.Error("Insufficient funds for debit", "err", balanceLow.Error())
		return balanceLow
	}

	err = s.store.DebitTransaction(ctx, t)
	if err != nil {
		s.logger.Error("Error in create debit transaction", "msg", err.Error(), req, t)
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
