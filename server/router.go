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

	transactionRoutes := router.PathPrefix("/transaction").Subrouter()
	transactionRoutes.Use(middleware.TransactionMiddleware)
	transactionRoutes.HandleFunc("/debit", transaction.DebitAmount(dep.TransactionService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	transactionRoutes.HandleFunc("/{account_id}", transaction.List(dep.TransactionService)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/createuser", user.Create(dep.UserServices)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/user/{user_id}", user.DeleteByID(dep.UserServices)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/users/{userId}", user.Update(dep.UserServices)).Methods(http.MethodPut)
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)

	//User
	router.HandleFunc("/createuser", middleware.AuthorizationMiddleware(user.Create(dep.UserServices), "createUser")).Methods(http.MethodPost).Headers(versionHeader, v1)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello")
}
