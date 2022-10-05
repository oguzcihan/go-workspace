package handlers

import (
	. "ExampleApp/internal/dtos/user"
	. "ExampleApp/internal/helpers"
	. "ExampleApp/internal/services"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type IUserHandler interface {
	//Create(w http.ResponseWriter, r *http.Request)
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

//func (userHandler UserHandler) Create(w http.ResponseWriter, r *http.Request) {
//	var user UserDto
//
//	jsonErr := json.NewDecoder(r.Body).Decode(&user)
//	if errors.As(jsonErr, &customError) {
//		//debug edilerek bakılacak
//		JSON(w, customError.Status, bodyError)
//		return
//	}
//	requestErr := RequestValidate(user)
//	if requestErr != nil {
//		//will work if validate
//		JSON(w, http.StatusBadRequest, requestErr)
//		return
//	}
//	//if true,registration start
//	resCreateUser, errs := userHandler.service.Create(user)
//	if errors.As(errs, &customError) {
//		//if an error occurs while recording
//		JSON(w, customError.Status, errs)
//		return
//	}
//
//	//Registration successful code send
//	JSON(w, http.StatusCreated, resCreateUser)
//
//}

func (userHandler UserHandler) Save(w http.ResponseWriter, r *http.Request) {
	//TODO update user
	var user UserDto

	Id := mux.Vars(r)["id"]
	if len(Id) == 0 {
		JSON(w, http.StatusBadRequest, emptyId)
		return
	}

	jsonErr := json.NewDecoder(r.Body).Decode(&user)
	if errors.As(jsonErr, &customError) {
		//debug edilerek bakılacak
		JSON(w, customError.Status, bodyError)
		return
	}

	validateResponse := RequestValidate(user)
	if validateResponse != nil {
		JSON(w, http.StatusBadRequest, validateResponse)
		return
	}

	convId, err := strconv.Atoi(Id)
	if err != nil {
		JSON(w, http.StatusBadRequest, notConverted)
		return
	}

	saveResponse := userHandler.service.Save(user, convId)
	if errors.As(saveResponse, &customError) {
		JSON(w, customError.Status, saveResponse)
		return
	}
}

func (userHandler UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]
	if len(Id) == 0 {
		JSON(w, http.StatusBadRequest, emptyId)
		return
	}

	convId, convErr := strconv.Atoi(Id)
	if convErr != nil {
		JSON(w, http.StatusBadRequest, notConverted)
		return
	}

	err := userHandler.service.Delete(convId)
	if errors.As(err, &customError) {
		JSON(w, customError.Status, err)
		return
	}

	JSON(w, http.StatusOK, err)

}

func (userHandler UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	//var userFilter filter.UserFilter
	generateUserPagination, paginationErr := GeneratePaginationRequest(r)
	if paginationErr != nil {
		JSON(w, http.StatusBadRequest, paginationErr)
	}
	//filterOptions, paginationOptions
	getResponse, err := userHandler.service.GetAllUser(generateUserPagination)
	if errors.As(err, &customError) {
		JSON(w, customError.Status, err)
		return
	}
	JSON(w, http.StatusOK, getResponse)
}
