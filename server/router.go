package server

import (
	"fmt"
	"net/http"

	"bankapp/config"
	"bankapp/login"
	"bankapp/middleware"
	"bankapp/transaction"
	"bankapp/user"
	"bankapp/useraccount"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	router = mux.NewRouter()
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/transaction/debit", middleware.AuthorizationMiddleware(transaction.Debit(dep.TransactionService))).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/transaction/credit", middleware.AuthorizationMiddleware(transaction.Credit(dep.TransactionService))).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/transaction/{account_id}", middleware.AuthorizationMiddleware(transaction.List(dep.TransactionService))).Methods(http.MethodGet).Headers(versionHeader, v1)
	// user
	router.HandleFunc("/createuseraccount", middleware.AuthorizationMiddleware(useraccount.Create(dep.UserAccountService))).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/user/{user_id}", middleware.AuthorizationMiddleware(user.DeleteByID(dep.UserServices))).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/users/{userId}", middleware.AuthorizationMiddleware(user.Update(dep.UserServices))).Methods(http.MethodPut)
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	return
}
