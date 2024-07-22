package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Define a test struct
type TestStruct struct {
	Name string
	Age  int
}

func TestUnPack(t *testing.T) {
	// Test case 1: Input as []byte
	t.Run("input as []byte", func(t *testing.T) {
		input := []byte(`{"Name":"Alice","Age":30}`)
		var target TestStruct

		err := UnPack(input, &target)
		assert.NoError(t, err)
		assert.Equal(t, "Alice", target.Name)
		assert.Equal(t, 30, target.Age)
	})

	// Test case 2: Input as struct
	t.Run("input as struct", func(t *testing.T) {
		input := TestStruct{Name: "Bob", Age: 25}
		var target TestStruct

		err := UnPack(input, &target)
		assert.NoError(t, err)
		assert.Equal(t, "Bob", target.Name)
		assert.Equal(t, 25, target.Age)
	})

	// Test case 3: Input as map
	t.Run("input as map", func(t *testing.T) {
		input := map[string]interface{}{
			"Name": "Charlie",
			"Age":  20,
		}
		var target TestStruct

		err := UnPack(input, &target)
		assert.NoError(t, err)
		assert.Equal(t, "Charlie", target.Name)
		assert.Equal(t, 20, target.Age)
	})

	// Test case 4: Invalid JSON
	t.Run("invalid JSON", func(t *testing.T) {
		input := []byte(`{"Name":"Alice","Age":}`) // Invalid JSON
		var target TestStruct

		err := UnPack(input, &target)
		assert.Error(t, err)
	})

	// Test case 5: Unmarshal into incompatible type
	t.Run("unmarshal into incompatible type", func(t *testing.T) {
		input := []byte(`{"Name":"Alice","Age":"thirty"}`) // Age is a string
		var target TestStruct

		err := UnPack(input, &target)
		assert.Error(t, err)
	})
}
