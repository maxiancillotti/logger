package logger

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func addFields(requestID string, fields ...zap.Field) []zap.Field {

	fields = append(fields,
		zap.Any("request_id", requestID),
	)

	return fields
}

func addErrorFields(err error, status int, fields ...zap.Field) []zap.Field {

	if err == nil {
		fields = append(fields,
			zap.Any("error", "error received is nil"),
			zap.Any("status", status),
		)
		return fields
	}

	errorMsg := err.Error()
	errorStack := fmt.Sprintf("%+v\n", err)
	errorCause := errors.Cause(err).Error()

	if errorStack == fmt.Sprintf("%v\n", err) {
		errorStack = ""
	}

	fields = append(fields,
		zap.Any("error", errorMsg),
		zap.Any("error_stack", errorStack),
		zap.Any("error_cause", errorCause),
		zap.Any("status", status),
	)

	return fields
}
