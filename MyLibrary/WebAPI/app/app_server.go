package app

import (
	"MyLibrary/WebAPI/roots"
	"MyLibrary/infrastructure"
)

func LoadApp() {
	infrastructure.AddInfrastructureService()
	roots.AllRoutes()
}
