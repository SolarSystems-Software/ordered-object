package orderedobject

import (
	"encoding/json"
	"strings"
)

type pair[V any] struct {
	key   string
	value V
}

// Object represents a JSON object that respects insertion order.
type Object[V any] []pair[V]

// Set sets key in object with the given value.
//
// The key is replaced if it already exists.
func (object *Object[V]) Set(key string, value V) {
	for _, pair := range *object {
		if pair.key == key {
			pair.value = value
			return
		}
	}

	*object = append(*object, pair[V]{key, value})
}

// Has reports if the given key is set.
func (object *Object[V]) Has(key string) bool {
	for _, pair := range *object {
		if pair.key == key {
			return true
		}
	}

	return false
}

// Get gets the value of key.
//
// The returned value is V's zero value if key isn't set.
func (object *Object[V]) Get(key string) V {
	for _, pair := range *object {
		if pair.key == key {
			return pair.value
		}
	}

	// "hack" to get V's zero value
	var empty V
	return empty
}

// MarshalJSON encodes the object into JSON format, respecting insertion order in the process.
func (object *Object[V]) MarshalJSON() ([]byte, error) {
	var builder strings.Builder

	// Start of object
	builder.WriteString("{")

	for _, pair := range *object {
		// Write comma if this isn't the first entry
		if builder.Len() > 1 {
			builder.WriteString(",")
		}

		// Write key
		encodedKey, err := json.Marshal(pair.key)
		if err != nil {
			return nil, err
		}
		builder.WriteString(string(encodedKey))

		// Write separator
		builder.WriteString(":")

		// Write value
		encodedValue, err := json.Marshal(pair.value)
		if err != nil {
			return nil, err
		}
		builder.WriteString(string(encodedValue))
	}

	// End of object
	builder.WriteString("}")

	return []byte(builder.String()), nil
}
