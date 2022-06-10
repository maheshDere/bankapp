package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/transaction"
	"bankapp/user"
	"fmt"
)

type dependencies struct {
	TransactionService transaction.Service
	UserServices       user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	transactionService := transaction.NewService(dbStore, logger)

	// call new service
	// remove println later

	userService := user.NewService(dbStore, logger)
	fmt.Println(logger, dbStore)
	return dependencies{
		TransactionService: transactionService,
		UserServices:       userService,
	}, nil
}
