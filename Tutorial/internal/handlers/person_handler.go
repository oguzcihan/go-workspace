package handlers

import (
	. "Tutorial/internal/models"
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
	var person Person
	ctx.BindJSON(&person)
	result := p.service.Create(&person)
	ctx.JSON(200, result)
}
