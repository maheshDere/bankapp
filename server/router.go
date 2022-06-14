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

	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	//Transaction routes
	router.HandleFunc("/transaction/debit", middleware.AuthorizationMiddleware(transaction.DebitAmount(dep.TransactionService), "customer")).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/transaction/{account_id}", middleware.AuthorizationMiddleware(transaction.FindByID(dep.TransactionService), "customer")).Methods(http.MethodGet).Headers(versionHeader, v1)
	//User routes
	router.HandleFunc("/user", middleware.AuthorizationMiddleware(user.Create(dep.UserServices), "accountant")).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/user/{user_id}", middleware.AuthorizationMiddleware(user.DeleteByID(dep.UserServices), "accountant")).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/user/{userId}", middleware.AuthorizationMiddleware(user.Update(dep.UserServices), "accountant")).Methods(http.MethodPut)
	return
}
