package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.SugaredLogger
var NamedLogger *zap.Logger

func getLogger(isLocal bool) *zap.SugaredLogger {
	var config zap.Config

	if isLocal {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "time"

	logger, _ := config.Build()
	defer logger.Sync()

	return logger.Sugar()
}

func getNamedLogger(isLocal bool) *zap.Logger {
	var config zap.Config

	if isLocal {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "time"

	logger, _ := config.Build()
	defer logger.Sync()

	return logger.Named("test")
}

func NewLoggers() {
	isLocal := os.Getenv("IS_LOCAL") != ""

	Logger = getLogger(isLocal)
	NamedLogger = getNamedLogger(isLocal)
}
