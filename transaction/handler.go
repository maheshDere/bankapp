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
