package logger

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func addFields(requestID string, fields ...zap.Field) []zap.Field {

	fields = append(fields,
		zap.String("request_id", requestID),
	)

	return fields
}

func addErrorFields(err error, status int, fields ...zap.Field) []zap.Field {

	if err == nil {
		fields = append(fields,
			zap.String("error", "error logged is nil"),
			zap.Int("status", status),
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
		zap.String("error", errorMsg),
		zap.String("error_stack", errorStack),
		zap.String("error_cause", errorCause),
		zap.Int("status", status),
	)

	return fields
}
