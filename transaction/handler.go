package transaction

import (
	"bankapp/api"
	"bankapp/app"
	"encoding/json"
	"net/http"
)

func Debit(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var d DebitCreditRequest
		err := json.NewDecoder(req.Body).Decode(&d)
		if err != nil {
			app.GetLogger().Warn("Error debit transaction", "msg", err.Error(), "transaction", req.Body)
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		balance, err := service.debitAmount(req.Context(), d)
		if isBadRequest(err) || err == invalidUserID || err == balanceLow {
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{
				Message: err.Error(),
			})
			return
		}

		api.Success(rw, http.StatusCreated, &CreateTransactionResponse{
			Message:      "Amount debited successfully",
			TotalBalance: balance,
		})
	})
}

func Credit(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var d DebitCreditRequest
		err := json.NewDecoder(req.Body).Decode(&d)
		if err != nil {
			app.GetLogger().Warn("Error credit transaction", "msg", err.Error(), "transaction", req.Body)
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		balance, err := service.creditAmount(req.Context(), d)
		if isBadRequest(err) || err == invalidUserID || err == balanceLow {
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{
				Message: err.Error(),
			})
			return
		}

		api.Success(rw, http.StatusCreated, &CreateTransactionResponse{
			Message:      "Amount credited successfully",
			TotalBalance: balance,
		})
	})
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var df ListRequest
		err := json.NewDecoder(req.Body).Decode(&df)
		if err != nil {
			app.GetLogger().Warn("Error listing transaction", "msg", err.Error(), "transaction", req.Body)
			api.Error(rw, http.StatusBadRequest, err.Error())
			return
		}

		resp, err := service.list(req.Context(), df)
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
