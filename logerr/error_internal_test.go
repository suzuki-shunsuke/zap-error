package logerr

import (
	"errors"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func Test_logError_Error(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		err  *logError
		exp  string
	}{
		{
			name: "nil",
		},
		{
			name: "no fields",
			err: &logError{
				err: errors.New("foo"),
			},
			exp: "foo",
		},
		{
			name: "fields",
			err: &logError{
				err: errors.New("foo"),
				fields: []zap.Field{
					zap.String("name", "yoo"),
				},
			},
			exp: "foo",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			s := d.err.Error()
			if s != d.exp {
				t.Fatalf("wanted %s, got %s", d.exp, s)
			}
		})
	}
}

func Test_logError_Unwrap(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		err  *logError
		exp  error
	}{
		{
			name: "nil",
		},
		{
			name: "normal",
			err: &logError{
				err: errors.New("foo"),
			},
			exp: errors.New("foo"),
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			e := d.err.Unwrap()
			if !reflect.DeepEqual(d.exp, e) {
				t.Fatalf("wanted %+v, got %+v", d.exp, e)
			}
		})
	}
}

func Test_logError_Fields(t *testing.T) { //nolint:funlen
	t.Parallel()
	data := []struct {
		name string
		err  *logError
		exp  []zap.Field
	}{
		{
			name: "nil",
		},
		{
			name: "normal",
			err: &logError{
				err: errors.New("foo"),
			},
			exp: nil,
		},
		{
			name: "normal 2",
			err: &logError{
				err: errors.New("foo"),
				fields: []zap.Field{
					zap.String("name", "foo"),
				},
			},
			exp: []zap.Field{
				zap.String("name", "foo"),
			},
		},
		{
			name: "normal 2",
			err: &logError{
				err: &logError{
					err: errors.New("foo"),
					fields: []zap.Field{
						zap.String("name", "foo"),
					},
				},
			},
			exp: []zap.Field{
				zap.String("name", "foo"),
			},
		},
		{
			name: "normal 3",
			err: &logError{
				err: &logError{
					err: errors.New("foo"),
					fields: []zap.Field{
						zap.String("name", "foo"),
					},
				},
				fields: []zap.Field{
					zap.String("id", "yoo"),
				},
			},
			exp: []zap.Field{
				zap.String("name", "foo"),
				zap.String("id", "yoo"),
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			fields := d.err.Fields()
			if !reflect.DeepEqual(d.exp, fields) {
				t.Fatalf("wanted %+v, got %+v", d.exp, fields)
			}
		})
	}
}
