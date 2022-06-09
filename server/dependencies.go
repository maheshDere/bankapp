package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/transaction"
	"fmt"
)

type dependencies struct {
	TransactionService transaction.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	transactionService := transaction.NewService(dbStore, logger)

	// call new service
	// remove println later
	fmt.Println(logger, dbStore)
	return dependencies{
		TransactionService: transactionService,
	}, nil
}
