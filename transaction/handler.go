package transaction

import (
	"bankapp/api"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DebitAmount(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var d debitCreditRequest
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
		}

		balance, err := service.debitAmount(r.Context(), d)
		if isBadRequest(err) || err == invalidUserID {
			api.Error(w, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		if err != nil {
			api.Error(w, http.StatusInternalServerError, api.Response{
				Message: err.Error(),
			})
			return
		}
		api.Success(w, http.StatusCreated, &createTransactionResponse{
			Message:      "Amount debited successfully",
			TotalBalance: balance,
		})
	})
}

func Credit(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var d debitCreditRequest
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
		}

		balance, err := service.creditAmount(r.Context(), d)
		if isBadRequest(err) || err == invalidUserID {
			api.Error(w, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		if err != nil {
			api.Error(w, http.StatusInternalServerError, api.Response{
				Message: err.Error(),
			})
			return
		}
		api.Success(w, http.StatusCreated, &createTransactionResponse{
			Message:      "Amount credited successfully",
			TotalBalance: balance,
		})
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Inside the FindByID handler")
		vars := mux.Vars(req)

		resp, err := service.findByID(req.Context(), vars["account_id"])

		if err == errNoAccountId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func isBadRequest(err error) bool {
	return err == invalidAmount
}
