package transaction

import (
	"bankapp/db"
	"context"

	"go.uber.org/zap"
)

type transactionService struct {
	store  db.TransactionStorer
	logger *zap.SugaredLogger
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
