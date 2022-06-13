package user

import (
	"bankapp/api"
	"encoding/json"
	"net/http"
)

func Create(service Service) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cr createRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
		}

		resp, err := service.create(req.Context(), cr)
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{
				Message: err.Error(),
			})
		}
		//add error handling code and check is req good or not
		api.Success(rw, http.StatusCreated, resp)
	})
}
