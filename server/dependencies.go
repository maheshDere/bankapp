package server

import (
	"bankapp/app"
	"bankapp/db"
	"bankapp/users"
	"fmt"
)

type dependencies struct {
	UserService users.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service
	userService := users.NewService(dbStore, logger)
	// remove println later
	fmt.Println(logger, dbStore)
	return dependencies{
		UserService: userService,
	}, nil
}
