package logerr_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/zap-error/logerr"
	"go.uber.org/zap"
)

func TestWithFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		err    error
		fields []zap.Field
		exp    error
	}{
		{
			name: "nil",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			err := logerr.WithFields(d.err, d.fields...)
			if !reflect.DeepEqual(d.exp, err) {
				t.Fatalf("wanted %+v, got %+v", d.exp, err)
			}
		})
	}
}

func TestToFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		err  error
		exp  []zap.Field
	}{
		{
			name: "nil",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			fields := logerr.ToFields(d.err)
			if !reflect.DeepEqual(d.exp, fields) {
				t.Fatalf("wanted %+v, got %+v", d.exp, fields)
			}
		})
	}
}
