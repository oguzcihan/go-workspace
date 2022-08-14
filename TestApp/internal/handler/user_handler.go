package handler

import (
	"TestApp/internal/dtos"
	. "TestApp/internal/service"
	"TestApp/internal/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

// TODO: TC algoritması kullanılarak gerçekten TC mi kontrolü yapılabilir
const (
	alreadyUsername = "Kullanıcı adı daha önce alınmış!"
	emptyBody       = "Boş değer gönderilemez"
	checkIdString   = "Id boş olamaz!"
	jsonKey         = "Content-Type"
	jsonValue       = "application/json"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(_service UserService) UserHandler {
	return &userHandler{service: _service}
}

type userHandler struct {
	service UserService
}

func (uHandler *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto
	bodyError := json.NewDecoder(r.Body).Decode(&user)
	//defer r.Body.Close()
	if bodyError != nil {
		//is body empty?
		http.Error(w, emptyBody, http.StatusBadRequest)
		return
	}
	requestErr := utils.RequestValidate(user)
	if requestErr != nil {
		//will work if validate
		err := JSON(w, http.StatusBadRequest, requestErr)
		if err != nil {
			return
		}
		return
	}
	//username unique check
	resUsername, checkUser := uHandler.CheckUserName(r, user.UserName)
	if resUsername != nil && checkUser {
		err := JSON(w, http.StatusBadRequest, resUsername)
		if err != nil {
			return
		}
	} else {
		//if true,registration start
		resCreateUser, errs := uHandler.service.Create(r.Context(), user)
		if errs != nil {
			//if an error occurs while recording
			err := JSON(w, http.StatusInternalServerError, resCreateUser)
			if err != nil {
				return
			}
			return
		}
		//Registration successful code send
		err := JSON(w, http.StatusCreated, resCreateUser)
		if err != nil {
			return
		}
	}

}

func (uHandler *userHandler) CheckUserName(r *http.Request, userName string) ([]string, bool) {
	//send to service layer for username
	var errorString []string
	resUsername, err := uHandler.service.GetUsername(r.Context(), userName)
	if err != nil {
		//Exit
		os.Exit(1)
		return nil, false
	} else if resUsername.UserName == userName {
		//Throw an error if the response from the DB and the response from the request are equal
		errorString = append(errorString, alreadyUsername)
		return errorString, true
	}
	return errorString, false
}

func (uHandler *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto
	bodyError := json.NewDecoder(r.Body).Decode(&user)
	if bodyError != nil {
		http.Error(w, emptyBody, http.StatusBadRequest)
		return
	}
	requestId := mux.Vars(r)["id"]
	if len(requestId) == 0 {
		http.Error(w, checkIdString, http.StatusBadRequest)
		return
	}
	idString := strconv.Itoa(user.ID)
	if idString == "0" {
		idString = requestId
	} else if requestId != idString {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	resUpdateUser, err := uHandler.service.Update(r.Context(), user)
	if err != nil {
		errJson := JSON(w, http.StatusBadRequest, resUpdateUser)
		if errJson != nil {
			return
		}
	}
	errJson := JSON(w, http.StatusOK, resUpdateUser)
	//http.Error(w, "Güncellendi", http.StatusOK)
	if errJson != nil {
		return
	}

}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set(jsonKey, jsonValue)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

/*
Marshall ve ioreader bellekten gelen verilerde kullanmak performansı daha iyi yapar
Body den alınan verilerde NewDecoder ve Decode daha performanslıdır.
*/
