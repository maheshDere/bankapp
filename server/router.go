package server

import (
	"fmt"

	"bankapp/config"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	// TODO: add doc
	// v2 := fmt.Sprintf("application/vnd.%s.v2", config.AppName())

	// Remove v1 later
	fmt.Println(v1)

	router = mux.NewRouter()
	return
}
