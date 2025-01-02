package logx

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nmarsollier/resourcesgo/internal/tools/env"
	"github.com/stretchr/testify/assert"
)

func TestNewFields(t *testing.T) {
	// Mock env.Get().ServerName
	originalEnvGet := env.Get
	env.Get = func() *env.Configuration {
		return &env.Configuration{ServerName: "test-server"}
	}
	defer func() { env.Get = originalEnvGet }()

	fields := NewFields()

	assert.Equal(t, "test-server", fields[SERVER])
	_, err := uuid.Parse(fields[THREAD])
	assert.NoError(t, err)
}

func TestAdd(t *testing.T) {
	fields := NewFields()
	fields = fields.Add("key", "value").Add("key2", "value2")

	assert.Equal(t, "value", fields["key"])
	assert.Equal(t, "value2", fields["key2"])
}
