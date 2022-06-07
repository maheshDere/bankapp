package server

import (
	"bankapp/app"
	"bankapp/db"
	"fmt"
)

type dependencies struct {
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service
	// remove println later
	fmt.Println(logger, dbStore)
	return dependencies{}, nil
}
