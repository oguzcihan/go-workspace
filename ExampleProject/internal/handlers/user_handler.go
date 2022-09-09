package handlers

import (
	"ExampleProject/internal/dtos"
	. "ExampleProject/internal/services"
	. "ExampleProject/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserHandler interface {
	Create(context *gin.Context)
	Save(context *gin.Context)
	Delete(context *gin.Context)
	GetAllUser(context *gin.Context)
	Patch(context *gin.Context)
}

func NewUserHandler(_service UserService) IUserHandler {
	return UserHandler{service: _service}
}

type UserHandler struct {
	service UserService
}

var (
	customError  CustomError
	emptyBody    = NewError("empty_body", http.StatusBadRequest)
	emptyId      = NewError("empty_request_id", http.StatusBadRequest)
	notConverted = NewError("parse_error", http.StatusBadRequest)
)

func (userHandler UserHandler) Create(context *gin.Context) {
	//TODO create user
	var userDto dtos.UserDto
	errJson := context.ShouldBindJSON(&userDto)
	if errJson != nil {
		JSON(context, http.StatusBadRequest, notConverted)
		return
	}
	validateResponse := RequestValidate(userDto)
	if validateResponse != nil {
		JSON(context, http.StatusBadRequest, validateResponse)
		return
	}
	createResponse, err := userHandler.service.Create(userDto)
	if errors.As(err, &customError) {
		JSON(context, customError.Status, err)
		return
	}
	//DRY
	JSON(context, http.StatusCreated, createResponse)
}

func (userHandler UserHandler) Save(context *gin.Context) {
	//TODO update user
	var userDto dtos.UserDto
	Id := context.Param("id")
	if len(Id) == 0 {
		JSON(context, http.StatusBadRequest, emptyId)
		return
	}

	errJson := context.ShouldBindJSON(&userDto)
	if errJson != nil {
		//error nasıl uygulanmalı?
		JSON(context, http.StatusBadRequest, emptyBody)
		return
	}

	validateResponse := RequestValidate(userDto)
	if validateResponse != nil {
		JSON(context, http.StatusBadRequest, validateResponse)
		return
	}

	convId, err := strconv.Atoi(Id)
	if err != nil {
		JSON(context, http.StatusBadRequest, notConverted)
		return
	}

	saveResponse, errorMessage := userHandler.service.Save(userDto, convId)
	if errors.As(errorMessage, &customError) {
		JSON(context, customError.Status, errorMessage)
		return
	}
	JSON(context, http.StatusOK, saveResponse)
}

func (userHandler UserHandler) Delete(context *gin.Context) {
	Id := context.Param("id")
	if len(Id) == 0 {
		JSON(context, http.StatusBadRequest, emptyId)
		return
	}

	convId, convErr := strconv.Atoi(Id)
	if convErr != nil {
		JSON(context, http.StatusBadRequest, notConverted)
		return
	}

	deleteResponse, err := userHandler.service.Delete(convId)
	if errors.As(err, &customError) {
		JSON(context, customError.Status, err)
		return
	}

	JSON(context, http.StatusOK, deleteResponse)

}

func (userHandler UserHandler) GetAllUser(context *gin.Context) {
	generateUserPagination := GeneratePaginationRequest(context)
	//filterOptions, paginationOptions
	getResponse, err := userHandler.service.GetAllUser(generateUserPagination)
	if errors.As(err, &customError) {
		JSON(context, customError.Status, err)
		return
	}
	JSON(context, http.StatusOK, getResponse)

}

func (userHandler UserHandler) Patch(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func JSON(context *gin.Context, code int, response interface{}) {
	context.JSON(code, response)
}

/*
Marshall ve ioreader bellekten gelen verilerde kullanmak performansı daha iyi yapar
Body den alınan verilerde NewDecoder ve Decode daha performanslıdır.
*/
