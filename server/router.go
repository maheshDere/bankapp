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
	router = mux.NewRouter()
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)

	//User
	router.HandleFunc("/createuser", middleware.AuthorizationMiddleware(user.Create(dep.UserServices), true)).Methods(http.MethodPost).Headers(versionHeader, v1)

	//Transaction
	transactionRoutes := router.PathPrefix("/transaction").Subrouter()
	transactionRoutes.Use(middleware.TransactionMiddleware)
	transactionRoutes.HandleFunc("/debit", transaction.DebitAmount(dep.TransactionService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	transactionRoutes.HandleFunc("/{account_id}", transaction.FindByID(dep.TransactionService)).Methods(http.MethodGet).Headers(versionHeader, v1)

	sh := http.StripPrefix("/docs/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/docs/").Handler(sh)
	return
}
