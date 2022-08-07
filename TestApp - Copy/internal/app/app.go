package app

import (
	. "TestApp/internal/handler"
	. "TestApp/internal/model"
	. "TestApp/internal/repository"
	. "TestApp/internal/services"
	"context"
	_ "github.com/lib/pq"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApplicationContext struct {
	User UserHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	//godotenv.Load(".env")
	//sqlDriver := os.Getenv("SQL_DRIVER")
	//connectionString := os.Getenv("CONNECTION_STRING")
	//NewApp dışarıda olmalı
	ormDB, err := gorm.Open(p.Open(conf.Sql.DataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ormDB.AutoMigrate(&User{})
	db, err := ormDB.DB()
	if err != nil {
		return nil, err
	}

	userRepository := NewUserRepository(ormDB)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	return &ApplicationContext{
		User: userHandler,
	}, nil
}
