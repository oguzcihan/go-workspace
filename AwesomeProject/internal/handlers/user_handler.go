package handlers

import (
	. "AwesomeProject/internal/dtos/userDtos"
	. "AwesomeProject/internal/helpers"
	. "AwesomeProject/internal/services"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	jsonKey   = "Content-Type"
	jsonValue = "application/json"
)

type IUserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(_service UserService) IUserHandler {
	return UserHandler{service: _service}
}

type UserHandler struct {
	service UserService
}

var (
	customError  CustomError
	bodyError    = NewError("body_error", http.StatusBadRequest)
	emptyId      = NewError("empty_request_id", http.StatusBadRequest)
	notConverted = NewError("parse_error", http.StatusBadRequest)
)

func (userHandler UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user UserDto
	code := http.StatusCreated

	jsonErr := json.NewDecoder(r.Body).Decode(&user)
	if errors.As(jsonErr, &customError) {
		//debug edilerek bakılacak
		JSON(w, customError.Status, bodyError)
		return
	}

	requestErr := RequestValidate(user)
	if requestErr != nil {
		//will work if validate
		JSON(w, http.StatusBadRequest, requestErr)
		return
	}
	//if true,registration start
	resCreateUser, errs := userHandler.service.Create(user)
	if errors.As(errs, &customError) {
		//if an error occurs while recording
		JSON(w, customError.Status, errs)
		return
	}

	//Registration successful code send
	JSON(w, code, resCreateUser)

}

func (userHandler UserHandler) Save(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (userHandler UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (userHandler UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func JSON(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set(jsonKey, jsonValue)
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {

	}
}
