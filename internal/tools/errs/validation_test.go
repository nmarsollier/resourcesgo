package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationField(t *testing.T) {
	validation := NewValidationField("field1", "error message")

	assert.Equal(t, 1, len(validation.Messages))
	assert.Equal(t, "field1", validation.Messages[0].Path)
	assert.Equal(t, "error message", validation.Messages[0].Message)
}

func TestNewValidation(t *testing.T) {
	validation := NewValidation()

	assert.Equal(t, 0, len(validation.Messages))
}

func TestValidation_Add(t *testing.T) {
	validation := NewValidation()
	validation.Add("field1", "error message 1")
	validation.Add("field2", "error message 2")

	assert.Equal(t, 2, len(validation.Messages))
	assert.Equal(t, "field1", validation.Messages[0].Path)
	assert.Equal(t, "error message 1", validation.Messages[0].Message)
	assert.Equal(t, "field2", validation.Messages[1].Path)
	assert.Equal(t, "error message 2", validation.Messages[1].Message)
}

func TestValidation_Size(t *testing.T) {
	validation := NewValidation()
	validation.Add("field1", "error message 1")

	assert.Equal(t, 1, validation.Size())
}

func TestValidation_Error(t *testing.T) {
	validation := NewValidation()
	validation.Add("field1", "error message 1")

	jsonStr := validation.Error()
	expected := `{"messages":[{"path":"field1","message":"error message 1"}]}`
	assert.JSONEq(t, expected, jsonStr)
}
