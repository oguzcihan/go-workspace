package services

import . "AwesomeProject/internal/repository"

func NewAccountService(_repository UserRepository) AccountService {
	return AccountService{repository: _repository}
}

type AccountService struct {
	repository UserRepository
}
