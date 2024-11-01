package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataStore(t *testing.T) {
	ds := NewDataStore()
	t.Run("PutAndGet", func(t *testing.T) {
		ds.Put("key1", "value1")
		ds.Put("key2", "value2")

		tests := []struct {
			key      string
			exists   bool
			expected string
		}{
			{"key1", true, "value1"},
			{"key2", true, "value2"},
			{"key3", false, ""},
		}

		for _, test := range tests {
			value, exists := ds.Get(test.key)
			assert.Equal(t, test.expected, value)
			assert.Equal(t, test.exists, exists)
		}

	})

	t.Run("GetListKeys", func(t *testing.T) {
		expectedKeys := []string{"key1", "key2"}
		keys := ds.GetListKeys()
		assert.ElementsMatch(t, expectedKeys, keys, "expected keys to match")
	})

	t.Run("Delete", func(t *testing.T) {
		tests := []struct {
			key      string
			expected bool
		}{
			{"key1", true},
			{"key3", false},
		}

		for _, test := range tests {
			isDeleted := ds.Delete(test.key)
			assert.Equal(t, test.expected, isDeleted)
		}
	})
}
