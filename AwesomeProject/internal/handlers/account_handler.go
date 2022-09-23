package handlers

import (
	. "AwesomeProject/internal/services"
	"net/http"
)

type IAccountHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

func NewAccountHandler(_service AccountService) IAccountHandler {
	return AccountHandler{service: _service}
}

type AccountHandler struct {
	service AccountService
}

func (a AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a AccountHandler) Register(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
