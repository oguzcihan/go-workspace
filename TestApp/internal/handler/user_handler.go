package handler

import (
	. "TestApp/internal/model"
	. "TestApp/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(_service UserService) UserHandler {
	return &userHandler{service: _service}
}

type userHandler struct {
	service UserService
}

func (uHandler *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	error1 := json.NewDecoder(r.Body).Decode(&user)
	//defer r.Body.Close()
	if error1 != nil {
		http.Error(w, error1.Error(), http.StatusBadRequest)
		return
	}

	resUser, error2 := uHandler.service.Create(r.Context(), &user)
	if error2 != nil {
		http.Error(w, error2.Error(), http.StatusInternalServerError)
		return
	}

	err := JSON(w, http.StatusCreated, resUser)
	if err != nil {
		return
	}

}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
