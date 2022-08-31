package services

import . "Tutorial/internal/repository"

func NewPersonService(_repository PersonRepository) *PersonService {
	return &PersonService{repository: _repository}
}

type PersonService struct {
	repository PersonRepository
}
