package handlers

import (
	"ExampleProject/internal/dtos"
	. "ExampleProject/internal/services"
	"ExampleProject/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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
	customError utils.CustomError
)

func (userHandler UserHandler) Create(context *gin.Context) {
	//TODO create user
	var userDto dtos.UserDto

	errJson := context.ShouldBindJSON(&userDto)

	if errJson != nil {
		response := utils.RequestValidate(userDto)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	response, err := userHandler.service.Create(userDto)
	if errors.As(err, &customError) {
		context.JSON(customError.Status, err)
		return
	}

	context.JSON(http.StatusCreated, response)
}

func (userHandler UserHandler) Save(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (userHandler UserHandler) Delete(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (userHandler UserHandler) GetAllUser(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (userHandler UserHandler) Patch(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
