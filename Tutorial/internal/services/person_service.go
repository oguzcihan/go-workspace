package services

import (
	. "Tutorial/internal/models"
	. "Tutorial/internal/repository"
)

func NewPersonService(_repository PersonRepository) *PersonService {
	return &PersonService{repository: _repository}
}

type PersonService struct {
	repository PersonRepository
}

func (service PersonService) Create(person *Person) *Person {
	service.repository.Create(person)
	return person
}
