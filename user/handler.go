package user

import (
	"bankapp/api"
	"encoding/json"
	"fmt"
	"net/http"
)

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
