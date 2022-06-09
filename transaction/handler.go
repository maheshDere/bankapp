package transaction

import (
	"bankapp/api"
	"encoding/json"
	"net/http"
)

func DebitAmount(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var d debitRequest
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
		}

		err = service.debitAmount(r.Context(), d)
		if isBadRequest(err) {
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
		api.Success(w, http.StatusCreated, api.Response{
			Message: "Amount debited successfully",
		})
	})
}

func isBadRequest(err error) bool {
	return err == invalidAmount
}
