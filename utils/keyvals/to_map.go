package keyvals

import "emperror.dev/errors/utils/keyval"

// ToMap creates a map of key-value pairs from a variadic key-value pair slice.
// Deprecated: use emperror.dev/errors/utils/keyval.ToMap.
func ToMap(kvs []interface{}) map[string]interface{} {
	return keyval.ToMap(kvs)
}
