package useraccount

import (
	"bankapp/api"
	"encoding/json"
	"net/http"
)

func Create(service Service) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cr createRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		//add error handling code here
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: "Invalid input data",
			})
			return
		}

		resp, err := service.create(req.Context(), cr)
		//add error handling code and check is req good or not
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{
				Message: "Error while creating user",
			})
			return
		}
		api.Success(rw, http.StatusCreated, resp)
	})
}
