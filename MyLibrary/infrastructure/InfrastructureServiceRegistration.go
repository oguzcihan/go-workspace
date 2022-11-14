package infrastructure

import "MyLibrary/infrastructure/zap_logger"

func AddInfrastructureService() {
	zap_logger.InitializeLogger()
}
