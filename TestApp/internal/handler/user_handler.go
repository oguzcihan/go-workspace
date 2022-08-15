package handler

import (
	"TestApp/internal/dtos"
	. "TestApp/internal/service"
	"TestApp/internal/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// TODO: TC algoritması kullanılarak gerçekten TC mi kontrolü yapılabilir
const (
	emptyBody = "empty_body"
	jsonKey   = "Content-Type"
	jsonValue = "application/json"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(_service UserService) UserHandler {
	return &userHandler{service: _service}
}

type userHandler struct {
	service UserService
}

func (uHandler *userHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user dtos.UserDto
	bodyError := json.NewDecoder(r.Body).Decode(&user)
	//defer r.Body.Close()
	if bodyError != nil {
		//is body empty?
		http.Error(w, emptyBody, http.StatusBadRequest)
		return
	}

	requestErr := utils.RequestValidate(user)
	if requestErr != nil {
		//will work if validate
		_ = JSON(w, http.StatusBadRequest, requestErr)
		return
	}

	//if true,registration start
	resCreateUser, errs := uHandler.service.Create(r.Context(), user)
	if errs != nil {
		//if an error occurs while recording
		_ = JSON(w, http.StatusInternalServerError, resCreateUser)
		return
	}

	//Registration successful code send
	_ = JSON(w, http.StatusCreated, resCreateUser)

}

func (uHandler *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto
	bodyError := json.NewDecoder(r.Body).Decode(&user)
	if bodyError != nil {
		http.Error(w, emptyBody, http.StatusBadRequest)
		return
	}

	requestId := mux.Vars(r)["id"]
	if len(requestId) == 0 {
		http.Error(w, "checkIdString", http.StatusBadRequest)
		return
	}
	//json içinden gelmemeli
	idString := strconv.Itoa(user.ID)
	if idString == "0" {
		idString = requestId
	} else if requestId != idString {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	resUpdateUser, err := uHandler.service.Update(r.Context(), user)
	if err != nil {
		errJson := JSON(w, http.StatusBadRequest, resUpdateUser)
		if errJson != nil {
			return
		}
	}
	errJson := JSON(w, http.StatusOK, resUpdateUser)
	//http.Error(w, "Güncellendi", http.StatusOK)
	if errJson != nil {
		return
	}

}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set(jsonKey, jsonValue)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

/*
Marshall ve ioreader bellekten gelen verilerde kullanmak performansı daha iyi yapar
Body den alınan verilerde NewDecoder ve Decode daha performanslıdır.
*/
