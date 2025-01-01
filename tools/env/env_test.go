package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Mock environment variables
	os.Setenv("SERVER_NAME", "test-server")
	os.Setenv("PORT", "8080")
	os.Setenv("GQL_PORT", "9090")
	os.Setenv("POSTGRES_URL", "postgres://test@localhost:5432/testdb")

	// Reset config to ensure it reloads
	config = nil

	cfg := Get()

	assert.Equal(t, "test-server", cfg.ServerName)
	assert.Equal(t, 8080, cfg.Port)
	assert.Equal(t, 9090, cfg.GqlPort)
	assert.Equal(t, "postgres://test@localhost:5432/testdb", cfg.PostgresURL)

	// Clean up environment variables
	os.Unsetenv("SERVER_NAME")
	os.Unsetenv("PORT")
	os.Unsetenv("GQL_PORT")
	os.Unsetenv("POSTGRES_URL")
}

func TestLoadDefaults(t *testing.T) {
	// Ensure environment variables are not set
	os.Unsetenv("SERVER_NAME")
	os.Unsetenv("PORT")
	os.Unsetenv("GQL_PORT")
	os.Unsetenv("POSTGRES_URL")

	cfg := load()

	assert.Equal(t, "resourcesgo", cfg.ServerName)
	assert.Equal(t, 3000, cfg.Port)
	assert.Equal(t, 4000, cfg.GqlPort)
	assert.Equal(t, "postgres://postgres@localhost:5432/postgres", cfg.PostgresURL)
}
