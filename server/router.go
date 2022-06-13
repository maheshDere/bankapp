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
	//JWT Authorization middleware
	router.Use(middleware.AuthorizationMiddleware)
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	//Transaction routes
	router.HandleFunc("/transaction/debit", transaction.DebitAmount(dep.TransactionService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/transaction/{account_id}", transaction.FindByID(dep.TransactionService)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//User routes
	router.HandleFunc("/user", user.Create(dep.UserServices)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/user/{user_id}", user.DeleteByID(dep.UserServices)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/user/{userId}", user.Update(dep.UserServices)).Methods(http.MethodPut)
	return
}
