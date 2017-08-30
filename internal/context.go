package internal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/goph/emperror"
)

// MapContext creates a map of key-value pairs.
//
// The implementation bellow is from go-kit's JSON logger.
func MapContext(err emperror.Contextor) map[string]interface{} {
	keyvals := err.Context()

	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd

	m := make(map[string]interface{}, n)

	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = errors.New("(MISSING)")

		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}

		merge(m, k, v)
	}

	return m
}

func merge(dst map[string]interface{}, k, v interface{}) {
	var key string

	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	switch x := v.(type) {
	case error:
		v = safeError(x)
	case fmt.Stringer:
		v = safeString(x)
	}

	dst[key] = v
}

func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(panicVal)
			}
		}
	}()

	s = str.String()

	return
}

func safeError(err error) (s interface{}) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				s = nil
			} else {
				panic(panicVal)
			}
		}
	}()

	s = err.Error()

	return
}
