package user

import (
	"bankapp/api"
	"bankapp/db"
	"encoding/json"
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

func DeleteByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := service.deleteByID(req.Context(), vars["user_id"])
		if err == db.ErrUserNotExist {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted Successfully"})
	})
}
