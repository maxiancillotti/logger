package logger

import "go.uber.org/zap"

type Logger interface {
	Flush()
	Field(key string, value interface{}) zap.Field

	Info(msg string, requestID string, fields ...zap.Field)
	Warn(msg string, requestID string, err error, status int, fields ...zap.Field)
	Error(msg string, requestID string, err error, status int, fields ...zap.Field)
	Debug(msg string, requestID string, err error, status int, fields ...zap.Field)
}

func New() Logger {
	return &globlalLogger
}
