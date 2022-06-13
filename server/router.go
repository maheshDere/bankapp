package server

import (
	"fmt"
	"net/http"

	"bankapp/config"
	"bankapp/middleware"
	"bankapp/transaction"
	"bankapp/users"

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

	transactionRoutes := router.PathPrefix("/transaction").Subrouter()
	transactionRoutes.Use(middleware.TransactionMiddleware)
	transactionRoutes.HandleFunc("/debit", transaction.DebitAmount(dep.TransactionService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	transactionRoutes.HandleFunc("/{account_id}", transaction.FindByID(dep.TransactionService)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/createuser", users.Create(dep.UserServices)).Methods(http.MethodPost).Headers(versionHeader, v1)

	router.HandleFunc("/users/{userId}", users.Update(dep.UserServices)).Methods(http.MethodPut)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}
