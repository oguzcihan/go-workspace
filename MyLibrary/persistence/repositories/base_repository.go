package repositories

import "MyLibrary/persistence/postgres"

var (
	postgresDb = postgres.ConnectDatabase()
)

func BaseRepositoryOperation() {
	NewUserRepository(postgresDb)
}
