package helper

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func InitNewLogger() *Logger {
	return &Logger{
		Logger: newLogger(),
	}
}

func newLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{
		"log/error.log",
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
