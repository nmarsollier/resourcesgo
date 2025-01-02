package logx

import (
	"github.com/google/uuid"
	"github.com/nmarsollier/resourcesgo/internal/tools/env"
)

const CORRELATION_ID = "correlation_id"
const CONTROLLER = "controller"
const RABBIT_ACTION = "rabbit_action"
const HTTP_METHOD = "http_method"
const HTTP_PATH = "http_path"
const HTTP_STATUS = "http_status"
const SERVER = "server"
const THREAD = "thread"

type Fields map[string]string

// NewFields creates a new Fields map with initial server and thread information.
func NewFields() Fields {
	return make(Fields, 4).
		Add(SERVER, env.Get().ServerName).
		Add(THREAD, uuid.New().String())
}

// Add inserts a key-value pair into the Fields map and returns the updated Fields.
//
// Parameters:
//   - key: The key to be added to the Fields map.
//   - value: The value associated with the key.
//
// Returns:
//   - Fields: The pointer to work as a builder.
func (f Fields) Add(key string, value string) Fields {
	f[key] = value
	return f
}
