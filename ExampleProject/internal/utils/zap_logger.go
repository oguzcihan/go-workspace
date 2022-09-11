package utils

import "go.uber.org/zap"

var (
	Logger *zap.Logger
)

func InitializeLogger() {
	var err error
	Logger, err = zap.NewProduction()
	//Logger, err = zap.NewDevelopment()
	if err != nil {
		Logger.Error("zap_error")
		return
	}

}
