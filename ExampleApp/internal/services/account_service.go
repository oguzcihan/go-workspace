package services

import (
	"ExampleApp/internal/dtos/account"
	. "ExampleApp/internal/helpers"
	"ExampleApp/internal/models"
	. "ExampleApp/internal/repository"
	"go.uber.org/zap"
	"time"
)

func NewAccountService(_repository UserRepository) AccountService {
	return AccountService{repository: _repository}
}

type AccountService struct {
	repository UserRepository
}

var (
	ErrorUserNameAlreadyExists  = NewError("username_already_exists", 400)
	ErrorUserEmailAlreadyExists = NewError("user_email_already_exists", 400)
	ErrorUserPassNotVerified    = NewError("user_password_could_not_be_verified", 400)
	ErrorTokenFailed            = NewError("failed_to_generate_token", 400)
	ErrorUserFailed             = NewError("user_not_found", 400)
	SuccessUserRegister         = NewError("success_user_register", 201)
)

func (accService AccountService) Register(user account.Register) error {
	/*
		repository e GetUsernameAndEmail kontrolü yapılması için veriler gönderilecek
		register içerisindeki password hash yapılacak
		en son create işlemi ele alınır
	*/

	err := accService.checkUserExistsForRegister(user.UserName, user.Email)
	if err != nil {
		Logger.Error("check_user_exists_error", zap.Error(err))
		return err
	}
	user.Password, err = GenerateHashPassword(user.Password)
	if err != nil {
		Logger.Error("password_hashing_error", zap.Error(err))
		return err
	}
	registerData := models.User{
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	_, dbErr := accService.repository.Create(&registerData)
	if dbErr != nil {
		return dbErr
	}
	return SuccessUserRegister
}

func (accService AccountService) Login(user account.Login) (*string, error) {
	var token account.Token

	userData, checkErr := accService.checkUserNameForLogin(user.Username)
	if checkErr != nil {
		return nil, checkErr
	}

	checkPass := CheckPasswordHash(user.Password, userData.Password)
	if !checkPass {
		return nil, ErrorUserPassNotVerified
	}

	verifiedToken, err := GenerateJWT(user.Username, userData.Role)
	if err != nil {
		return nil, ErrorTokenFailed
	}
	token.UserName = userData.UserName
	token.Role = userData.Role
	token.TokenString = *verifiedToken

	return verifiedToken, nil

}

func (accService AccountService) checkUserExistsForRegister(name string, email string) error {
	userData, err := accService.repository.GetUsernameAndEmail(name, email)
	if err != nil {
		return err
	}
	if userData.UserName == name {
		return ErrorUserNameAlreadyExists
	} else if userData.Email == email {
		return ErrorUserEmailAlreadyExists
	}
	return nil
}

func (accService AccountService) checkUserNameForLogin(userName string) (*models.User, error) {
	resUsername, _ := accService.repository.GetUsername(userName)
	if resUsername.UserName != userName {
		return nil, ErrorUserFailed
	}
	return resUsername, nil

}
