package handlers

import (
	"ExampleApp/internal/dtos/account"
	. "ExampleApp/internal/helpers"
	. "ExampleApp/internal/middleware"
	. "ExampleApp/internal/services"
	"encoding/json"
	"errors"
	"net/http"
)

type IAccountHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	AdminIndex(w http.ResponseWriter, r *http.Request)
	UserIndex(w http.ResponseWriter, r *http.Request)
}

func NewAccountHandler(_service AccountService) IAccountHandler {
	return AccountHandler{service: _service}
}

type AccountHandler struct {
	service AccountService
}

var (
	SuccessAdminPage = NewError("welcome_admin_page", http.StatusOK)
	SuccessUserPage  = NewError("welcome_user_page", http.StatusOK)
	SuccessIndexPage = NewError("welcome_public_page", http.StatusOK)
)

func (accountHandler AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser account.Login
	jsonErr := json.NewDecoder(r.Body).Decode(&loginUser)
	if errors.As(jsonErr, &customError) {
		//debug edilerek bakılacak
		JSON(w, customError.Status, bodyError)
		return
	}

	requestErr := RequestValidate(loginUser)
	if requestErr != nil {
		//will work if validate
		JSON(w, http.StatusBadRequest, requestErr)
		return
	}

	resLogin, err := accountHandler.service.Login(loginUser)
	if errors.As(err, &customError) {
		//if an error occurs while recording
		JSON(w, customError.Status, err)
		return
	}

	JSON(w, http.StatusOK, resLogin)

}

func (accountHandler AccountHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user account.Register

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
	userRegister := accountHandler.service.Register(user)
	if errors.As(userRegister, &customError) {
		//if an error occurs while recording
		JSON(w, customError.Status, userRegister)
	}
}

func (accountHandler AccountHandler) Index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Public Index Page"))
	JSON(w, http.StatusOK, SuccessIndexPage)
}

func (accountHandler AccountHandler) AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		JSON(w, http.StatusUnauthorized, ErrorNotAuthorized)
		return
	}
	JSON(w, http.StatusUnauthorized, SuccessAdminPage)
}

func (accountHandler AccountHandler) UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		JSON(w, http.StatusUnauthorized, ErrorNotAuthorized)
	}
	JSON(w, http.StatusOK, SuccessUserPage)
}
