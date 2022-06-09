package login

import (
	"bankapp/api"
	"encoding/json"
	"net/http"
)

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var u userLoginRequest
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.login(req.Context(), u)

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Login Successfully!!"})
	})
}
