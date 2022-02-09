package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger = getLogger()
var NamedLogger = getNamedLogger()

func getLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "time"

	logger, _ := config.Build()
	defer logger.Sync()

	return logger.Sugar()
}

func getNamedLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "time"

	logger, _ := config.Build()
	defer logger.Sync()

	return logger.Named("test")
}
