package zap_logger

import "go.uber.org/zap"

var Logger *zap.Logger

func InitializeLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		Logger.Error("base_zap_error")
		return
	}

}
