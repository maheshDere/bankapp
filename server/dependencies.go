package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/transaction"
)

type dependencies struct {
	TransactionService transaction.TransactionService
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	transactionService := transaction.NewService(dbStore, logger)

	return dependencies{
		TransactionService: transactionService,
	}, nil
}
