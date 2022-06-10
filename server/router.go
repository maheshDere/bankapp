package server

import (
	"bankapp/config"
	"bankapp/users"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	// TODO: add doc
	// v2 := fmt.Sprintf("application/vnd.%s.v2", config.AppName())
	fmt.Println(v1)
	router = mux.NewRouter()
	router.HandleFunc("/users/{userId}", users.Update(dep.UserService)).Methods(http.MethodPut)
	return
}
