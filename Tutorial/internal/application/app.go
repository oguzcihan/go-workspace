package application

import (
	"Tutorial/internal/config"
	"Tutorial/internal/handlers"
	. "Tutorial/internal/models"
	"Tutorial/internal/repository"
	"Tutorial/internal/services"
)

type ApplicationContext struct {
	PersonContext handlers.PersonHandler
}

func NewApplication() (*ApplicationContext, error) {

	database := config.DatabaseConnection(&Person{})

	personRepository := repository.NewPersonRepository(database)
	personService := services.NewPersonService(*personRepository)
	personHandler := handlers.NewPersonHandler(*personService)

	return &ApplicationContext{PersonContext: personHandler}, nil
}
