package logerr

import (
	"errors"

	"go.uber.org/zap"
)

func WithFields(err error, fields ...zap.Field) error {
	if err == nil {
		return nil
	}
	return &logError{
		err:    err,
		fields: fields,
	}
}

func ToFields(err error) []zap.Field {
	var e *logError
	if errors.As(err, &e) {
		return e.Fields()
	}
	return nil
}
