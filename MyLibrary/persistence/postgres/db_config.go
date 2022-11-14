package postgres

import (
	zaplog "MyLibrary/infrastructure/zap_logger"
	"go.uber.org/zap"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func ConnectDatabase() *gorm.DB {
	pgConnection := os.Getenv("PG_CONNECTION")
	connection, err := gorm.Open(pg.Open(pgConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zaplog.Logger.Fatal("database_url_error", zap.Error(err))
		return nil
	}

	sqlDB, errConnection := connection.DB()
	errConnection = sqlDB.Ping()
	if errConnection != nil {
		zaplog.Logger.Fatal("connection_error", zap.Error(errConnection))
	}

	zaplog.Logger.Info("database_connection_success")

	return connection
}
