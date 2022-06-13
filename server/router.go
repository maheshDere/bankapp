package server

import (
	"fmt"
	"net/http"

	"bankapp/config"
	"bankapp/login"
	"bankapp/middleware"
	"bankapp/transaction"
	"bankapp/user"

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
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)

	subrouter := router.PathPrefix("/").Subrouter()
	subrouter.Use(middleware.AuthorizationMiddleware)
	subrouter.HandleFunc("/transaction/debit", transaction.DebitAmount(dep.TransactionService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	subrouter.HandleFunc("/transaction/{account_id}", transaction.FindByID(dep.TransactionService)).Methods(http.MethodGet).Headers(versionHeader, v1)
	subrouter.HandleFunc("/user", user.Create(dep.UserServices)).Methods(http.MethodPost).Headers(versionHeader, v1)
	subrouter.HandleFunc("/user/{user_id}", user.DeleteByID(dep.UserServices)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	subrouter.HandleFunc("/user/{userId}", user.Update(dep.UserServices)).Methods(http.MethodPut)
	//Login

	//User
	subrouter.HandleFunc("/createuser", user.Create(dep.UserServices)).Methods(http.MethodPost).Headers(versionHeader, v1)
	return
}
