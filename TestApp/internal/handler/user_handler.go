package handler

import (
	"TestApp/internal/dtos"
	. "TestApp/internal/service"
	"TestApp/internal/utils"
	"encoding/json"
	"net/http"
)

// TODO: TC algoritması kullanılarak gerçekten TC mi kontrolü yapılabilir

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
	var user dtos.UserDto
	//
	//error1 := json.NewDecoder(r.Body).Decode(&user)
	//if error1 != nil {
	//	http.Error(w, error1.Error(), http.StatusBadRequest)
	//	return
	//}
	requestErr := utils.RequestValidate(user)
	if requestErr != nil {
		err := JSON(w, http.StatusBadRequest, requestErr)
		if err != nil {
			return
		}
		return
	}
	resUser, errs := uHandler.service.Create(r.Context(), user)
	if errs != nil {
		err := JSON(w, http.StatusInternalServerError, resUser)
		if err != nil {
			return
		}
		return
	} else {
		err := JSON(w, http.StatusCreated, resUser)
		if err != nil {
			return
		}
	}

}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
