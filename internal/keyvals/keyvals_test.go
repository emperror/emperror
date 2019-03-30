package keyvals

import (
	"testing"

	"github.com/pkg/errors"
)

func TestToKeyvals(t *testing.T) {
	m := ToMap([]interface{}{"key", "value", "error", errors.New("error")})

	expected := map[string]interface{}{
		"key":   "value",
		"error": "error",
	}

	if len(expected) != len(m) {
		t.Fatalf("expected log fields to be equal\ngot:  %v\nwant: %v", m, expected)
	}

	for key, value := range expected {
		if m[key] != value {
			t.Fatalf("expected log fields to be equal\ngot:  %v\nwant: %v", m, expected)
		}
	}
}
