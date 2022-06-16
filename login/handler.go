package login

import (
	"bankapp/api"
	"encoding/json"
	"net/http"
)

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var u LoginRequest
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		tokenString, err := service.login(req.Context(), u)
		if err != nil {
			api.Error(rw, http.StatusUnauthorized, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusCreated, api.Response{Message: tokenString})
	})
}
