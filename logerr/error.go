package logerr

import (
	"go.uber.org/zap"
)

type logError struct {
	err    error
	fields []zap.Field
}

func (e *logError) Error() string {
	if e == nil || e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *logError) Unwrap() error {
	if e == nil || e.err == nil {
		return nil
	}
	return e.err
}

func (e *logError) Fields() []zap.Field {
	if e == nil {
		return nil
	}
	return append(toFields(e.err), e.fields...)
}
