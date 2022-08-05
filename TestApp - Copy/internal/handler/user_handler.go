package handler

import (
	. "TestApp/internal/model"
	. "TestApp/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(service UserService) UserHandler {
	return &userHandler{service: service}
}

type userHandler struct {
	service UserService
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	error1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if error1 != nil {
		http.Error(w, error1.Error(), http.StatusBadRequest)
		return
	}

	res, error2 := h.service.Create(r.Context(), &user)
	if error2 != nil {
		http.Error(w, error2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	error1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if error1 != nil {
		http.Error(w, error1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id boş olamaz", http.StatusBadRequest)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id ler eşleşmedi", http.StatusBadRequest)
		return
	}

	res, error2 := h.service.Update(r.Context(), &user)
	if error2 != nil {
		http.Error(w, error2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "ID boş olamaz", http.StatusBadRequest)
		return
	}
	res, error := h.service.Delete(r.Context(), id)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
