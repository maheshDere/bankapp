package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/login"
	"bankapp/transaction"
	"bankapp/user"
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
	dbUserStore := db.NewStorer(appDB)
	dbLoginStore := db.NewLoginStorer(appDB)

	userService := user.NewService(dbUserStore, logger)
	loginService := login.NewService(dbLoginStore, logger)
	transactionService := transaction.NewService(dbStore, logger)

	return dependencies{
		UserLoginService:   loginService,
		TransactionService: transactionService,
		UserServices:       userService,
	}, nil
}
