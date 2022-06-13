package users

import (
	"bankapp/api"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Update(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c updateRequest
		userId := mux.Vars(req)["userId"]
		if userId == "" {
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: "Invalid user id",
			})
			return
		}

		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.update(req.Context(), c, userId)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
	})

}
func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyPassword
}

func Create(service Service) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cr createRequest
		_ = json.NewDecoder(req.Body).Decode(&cr)
		//add error handling code here

		fmt.Println("In create cr is --> ", cr)
		_ = service.create(req.Context(), cr)

		//add error handling code and check is req good or not
		api.Success(rw, http.StatusCreated, api.Response{Message: "User created sucessfully"})
	})
}
