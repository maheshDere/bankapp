package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/user"
	"fmt"
)

type dependencies struct {
	UserServices user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service
	// remove println later

	userService := user.NewService(dbStore, logger)
	fmt.Println(logger, dbStore)
	return dependencies{
		UserServices: userService,
	}, nil
}
