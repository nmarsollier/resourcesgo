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

func NewFields() Fields {
	return make(Fields, 4).
		Add(SERVER, env.Get().ServerName).
		Add(THREAD, uuid.New().String())
}

func (f Fields) Add(key string, value string) Fields {
	f[key] = value
	return f
}
