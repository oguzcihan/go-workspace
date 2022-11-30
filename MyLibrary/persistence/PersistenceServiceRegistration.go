package persistence

import (
	"MyLibrary/domain/entities"
	"MyLibrary/persistence/dotenv"
	"MyLibrary/persistence/migration"
	"MyLibrary/persistence/repositories"
)

func AddPersistenceService() {
	dotenv.LoadDotEnvFile()
	migration.InitialMigration(entities.User{})

	repositories.BaseRepositoryOperation()
}
