package handlers

import (
	. "Tutorial/internal/services"
	"github.com/gin-gonic/gin"
)

type PersonHandler interface {
	Create(ctx *gin.Context)
}

func NewPersonHandler(_service PersonService) PersonHandler {
	return &personHandler{service: _service}
}

type personHandler struct {
	service PersonService
}

func (p personHandler) Create(ctx *gin.Context) {

}
