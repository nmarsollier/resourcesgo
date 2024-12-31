package logx

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nmarsollier/resourcesgo/tools/env"
	"github.com/nmarsollier/resourcesgo/tools/strs"
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

func ErrorStr(fields Fields, err string) {
	logWithFields("ERROR", fields, err)
}

func Error(fields Fields, err error) {
	logWithFields("ERROR", fields, strs.ToJson(err))
}

func Info(fields Fields, msg string) {
	logWithFields("INFO", fields, msg)
}

func Warn(fields Fields, msg string) {
	logWithFields("WARN", fields, msg)
}

func logWithFields(level string, fields Fields, msg string) {
	logger := log.New(os.Stdout, fmt.Sprintf("%s: ", level), log.Ldate|log.Ltime|log.Lshortfile)
	data := ""
	for k, v := range fields {
		data += fmt.Sprintf("%s=%v ", k, v)
	}
	logger.Println(data + msg)
}
