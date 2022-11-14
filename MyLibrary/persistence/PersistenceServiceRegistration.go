package persistence

import (
	"MyLibrary/domain/entities"
	"MyLibrary/persistence/dotenv"
	"MyLibrary/persistence/migration"
)

func AddPersistenceService() {
	dotenv.LoadDotEnvFile()
	migration.InitialMigration(entities.User{})
}
