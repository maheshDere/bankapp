package transaction

import (
	"bankapp/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func FindByID(service TransactionService) http.HandlerFunc {
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
