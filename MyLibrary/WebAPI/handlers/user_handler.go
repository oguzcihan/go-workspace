package handlers

import (
	"MyLibrary/WebAPI/utils"
	"MyLibrary/application/features/auths/service"
	"MyLibrary/domain/entities"
	"MyLibrary/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService service.IUserService
	redisAuth   auth.IAuthentication
	tokenAuth   auth.IToken
}

// create users constructor
func NewUserHandler(_userService service.IUserService, redisAuth auth.IAuthentication, _tokenAuth auth.IToken) *UserHandler {
	return &UserHandler{
		userService: _userService,
		redisAuth:   redisAuth,
		tokenAuth:   _tokenAuth,
	}
}

func (uHandler UserHandler) Register(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	validateErr := utils.RequestValidate(user)
	if validateErr != nil {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}

	newUser, err := uHandler.userService.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newUser.PublicUser())
}
