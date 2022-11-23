package app

import (
	"MyLibrary/WebAPI/roots"
	"MyLibrary/application"
	"MyLibrary/infrastructure"
	"MyLibrary/persistence"
)

func StartApp() {
	infrastructure.AddInfrastructureService()
	persistence.AddPersistenceService()
	application.AddApplicationService()

	roots.AllRoutes()
}
