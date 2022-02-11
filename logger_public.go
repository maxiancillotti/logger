package logger

import "go.uber.org/zap"

type Logger interface {
	Info(msg string, requestID string, fields ...zap.Field)
	Warn(msg string, requestID string, err error, status int, fields ...zap.Field)
	Error(msg string, requestID string, err error, status int, fields ...zap.Field)
	Debug(msg string, requestID string, err error, status int, fields ...zap.Field)

	Flush()

	Field(key string, value interface{}) zap.Field
	StringField(key string, value string) zap.Field
	IntField(key string, value int) zap.Field
}

func New() Logger {
	build()
	return &globlalLogger
}
