package repositories

import (
	"MyLibrary/domain/repositories"
	"MyLibrary/infrastructure/persistence/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	User     repositories.IUserRepository
	database *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	dbConnection := postgres.ConnectDatabase()

	return &Repositories{
		User:     NewUserRepository(dbConnection),
		database: dbConnection,
	}, nil
}
