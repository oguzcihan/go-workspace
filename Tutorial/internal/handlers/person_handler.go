package handlers

import (
	. "Tutorial/internal/models"
	. "Tutorial/internal/services"
	"Tutorial/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PersonHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
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

func (p personHandler) GetAll(ctx *gin.Context) {
	code := http.StatusOK
	pagination := utils.GeneratePaginationRequest(ctx)
	response := p.service.GetAll(ctx, pagination)

	//if !response.Success {
	//	code = http.StatusBadRequest
	//}

	ctx.JSON(code, response)

}
