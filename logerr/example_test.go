package logerr_test

import (
	"errors"
	"fmt"

	"github.com/suzuki-shunsuke/zap-error/logerr"
	"go.uber.org/zap"
)

func Example() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck
	logger = logger.With(zap.String("program", "example"))
	if err := core(); err != nil {
		logger.Error("program exits", logerr.ToFields(err)...)
	}
}

func core() error {
	return fmt.Errorf("foo4: %w", foo4())
}

func foo1() error {
	return errors.New("foo")
}

func foo2() error {
	// Add fields to error
	return logerr.WithFields(foo1(), zap.String("name2", "foo2"), zap.String("foo2", "foo2"))
}

func foo3() error {
	return fmt.Errorf("foo2: %w", foo2())
}

func foo4() error {
	fields := []zap.Field{
		zap.String("name4", "foo4"),
	}
	// Add fields to error
	return logerr.WithFields(foo3(), append(fields, zap.String("foo4", "foo4"))...)
}
