package server

import (
	"bankapp/config"
	"bankapp/user"
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

	// Remove v1 later
	fmt.Println(v1)

	router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/createuser", user.Create(dep.UserServices)).Methods(http.MethodPost).Headers(versionHeader, v1)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}
