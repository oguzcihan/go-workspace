package services

import (
	. "TestApp/internal/model"
	. "TestApp/internal/repository"
	"context"
)

type UserService interface {
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

// Constructor
func NewUserService(repository UserRepository) UserService {
	return &userService{repository: repository}
}

type userService struct {
	repository UserRepository
}

// Create repository e context ve user gonderir
func (s *userService) Create(ctx context.Context, user *User) (int64, error) {
	//TODO console log yazdırılmalı
	return s.repository.Create(ctx, user)
}

// Update repository e update etmek için context ve user gönderir
func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	return s.repository.Update(ctx, user)
}

// Delete repository e delete yapmak için context ve id gönderir
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
