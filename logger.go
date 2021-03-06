package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	corelogger *zap.Logger
	once       sync.Once
}

var (
	globlalLogger logger // *zap.Logger
)

func build() {
	globlalLogger.once.Do(func() {
		logConfig := zap.Config{
			OutputPaths: []string{"stdout"},
			Encoding:    "json",
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "msg",
				LevelKey:     "level",
				TimeKey:      "time",
				EncodeTime:   zapcore.ISO8601TimeEncoder,
				EncodeLevel:  zapcore.LowercaseLevelEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}
		var err error
		globlalLogger.corelogger, err = logConfig.Build()
		if err != nil {
			panic(err)
		}
	})
}

// Flushes any buffered log entries. Applications should take care to call Flush before exiting.
func (l *logger) Flush() {
	l.corelogger.Sync()
}

// Logging methods

func (l *logger) Info(msg string, requestID string, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	l.corelogger.Info(msg, fields...)
}

func (l *logger) Warn(msg string, requestID string, err error, status int, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	fields = addErrorFields(err, status, fields...)
	l.corelogger.Warn(msg, fields...)
}

func (l *logger) Error(msg string, requestID string, err error, status int, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	fields = addErrorFields(err, status, fields...)
	l.corelogger.Error(msg, fields...)
}

func (l *logger) Debug(msg string, requestID string, err error, status int, fields ...zap.Field) {

	fields = addFields(requestID, fields...)
	fields = addErrorFields(err, status, fields...)
	l.corelogger.Debug(msg, fields...)
}

// Field methods

func (l *logger) Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func (l *logger) StringField(key string, value string) zap.Field {
	return zap.String(key, value)
}

func (l *logger) IntField(key string, value int) zap.Field {
	return zap.Int(key, value)
}
