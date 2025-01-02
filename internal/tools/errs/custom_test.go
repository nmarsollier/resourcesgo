package errs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustom(t *testing.T) {
	status := 404
	message := "Not Found"
	customErr := NewCustom(status, message)

	assert.Equal(t, status, customErr.Status())
	assert.Equal(t, message, customErr.Message)
}

func TestCustom_Error(t *testing.T) {
	status := 500
	message := "Internal Server Error"
	customErr := NewCustom(status, message)

	jsonStr := customErr.Error()
	expected := `"Internal Server Error"`
	assert.JSONEq(t, expected, jsonStr)
}

func TestCustom_Status(t *testing.T) {
	status := 400
	message := "Bad Request"
	customErr := NewCustom(status, message)

	assert.Equal(t, status, customErr.Status())
}
