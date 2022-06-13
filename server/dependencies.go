package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/login"
	"bankapp/transaction"
	"bankapp/user"
	"fmt"
)

type dependencies struct {
	UserLoginService   login.Service
	TransactionService transaction.Service
	UserServices       user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	transactionService := transaction.NewService(dbStore, logger)

	// call new service
	userService := user.NewService(dbStore, logger)

	loginService := login.NewService(dbStore, logger)

	// remove println later

	fmt.Println(logger, dbStore)
	return dependencies{
		TransactionService: transactionService,
		UserServices:       userService,
		UserLoginService:   loginService,
	}, nil
}
