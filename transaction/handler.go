package transaction

import (
	"bankapp/api"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Debit(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var d debitCreditRequest
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
		}

		balance, err := service.debitAmount(r.Context(), d)
		if isBadRequest(err) || err == invalidUserID || err == balanceLow {
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
		if isBadRequest(err) || err == invalidUserID || err == balanceLow {
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

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var df listRequest
		err := json.NewDecoder(req.Body).Decode(&df)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, err.Error())
			return
		}

		accountId := mux.Vars(req)["account_id"]
		resp, err := service.list(req.Context(), accountId, df)

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
