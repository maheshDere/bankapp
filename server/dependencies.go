package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/login"
	"fmt"
)

type dependencies struct {
	UserLoginService login.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)
	dbUserStore := db.NewLoginStorer(appDB)

	loginService := login.NewService(dbUserStore, logger)

	fmt.Println(logger, dbStore)
	return dependencies{
		UserLoginService: loginService,
	}, nil
}
