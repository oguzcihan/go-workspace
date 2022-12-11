package repositories

import (
	"MyLibrary/domain/entities"
)

type IUserRepository interface {
	Create(*entities.User) (*entities.User, error)
}
