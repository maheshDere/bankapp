package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/transaction"
	"bankapp/user"
	"bankapp/users"
	"fmt"
)

type dependencies struct {
	TransactionService transaction.Service
	UserServices       user.Service
	UserService        users.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	transactionService := transaction.NewService(dbStore, logger)

	// call new service
	userService := users.NewService(dbStore, logger)
	// remove println later

	fmt.Println(logger, dbStore)
	return dependencies{
		TransactionService: transactionService,
		UserServices:       userService,
	}, nil
}
