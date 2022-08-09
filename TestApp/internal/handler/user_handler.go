package handler

import (
	. "TestApp/internal/model"
	. "TestApp/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(_service UserService) UserHandler {
	return &userHandler{service: _service}
}

type userHandler struct {
	service UserService
}

func (uHandler *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	error1 := json.NewDecoder(r.Body).Decode(&user)
	//defer r.Body.Close()
	if error1 != nil {
		http.Error(w, error1.Error(), http.StatusBadRequest)
		return
	}
	//TODO: TC algoritması kullanılarak gerçekten TC mi kontrolü yapılabilir
	//checkTcNo(user.TcNo)
	//checkUserName(user.UserName)
	resUser, error2 := uHandler.service.Create(r.Context(), &user)
	if error2 != nil {
		http.Error(w, error2.Error(), http.StatusInternalServerError)
		return
	}

	err := JSON(w, http.StatusCreated, resUser)
	if err != nil {
		return

	}

}

func checkTcNo(tcNo int64) bool {
	lengthTcNo := len(strconv.Itoa(int(tcNo)))
	if lengthTcNo < 11 || lengthTcNo > 11 || lengthTcNo == 0 {
		return false
	} else {
		return true
	}
}
func checkUserName(userName string) string {
	return "nil"
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
