package logerr

import (
	"errors"

	"go.uber.org/zap"
)

// WithFields appends fields to err and returns a new error. If err is nil, nil is returned.
func WithFields(err error, fields ...zap.Field) error {
	if err == nil {
		return nil
	}
	return &logError{
		err:    err,
		fields: fields,
	}
}

// ToFields converts err to a list of zap.Field.
func ToFields(err error) []zap.Field {
	if err == nil {
		return nil
	}
	var e *logError
	if errors.As(err, &e) {
		return e.Fields()
	}
	return []zap.Field{zap.Error(err)}
}
