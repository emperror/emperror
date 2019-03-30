package keyvals_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/goph/emperror/internal/keyvals"
)

func TestToKeyvals(t *testing.T) {
	m := keyvals.ToMap([]interface{}{"key", "value", "error", errors.New("error")})

	assert.Equal(
		t,
		map[string]interface{}{
			"key":   "value",
			"error": "error",
		},
		m,
	)
}
