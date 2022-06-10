package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/user"
	"fmt"
)

type dependencies struct {
	UserService user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service
	// remove println later
	fmt.Println(logger, dbStore)

	userService := user.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
	}, nil
}
