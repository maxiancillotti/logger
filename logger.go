package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	logger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

// Flushes any buffered log entries. Applications should take care to call Flush before exiting.
func Flush() {
	logger.Sync()
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Debug(msg string, requestID string, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	logger.Debug(msg, fields...)
}

func Info(msg string, requestID string, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	logger.Info(msg, fields...)
}

func Warn(msg string, requestID string, err error, status int, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	fields = addErrorFields(err, status, fields...)
	logger.Error(msg, fields...)
}

func Error(msg string, requestID string, err error, status int, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	fields = addErrorFields(err, status, fields...)
	logger.Error(msg, fields...)
}
