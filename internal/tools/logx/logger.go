package logx

import (
	"fmt"
	"log"
	"os"
)

// Logger factory
var newLogger = func() *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func ErrorStr(fields Fields, err string) {
	logWithFields("ERROR", fields, err)
}

func Error(fields Fields, err error) {
	logWithFields("ERROR", fields, err.Error())
}

func Info(fields Fields, msg string) {
	logWithFields("INFO", fields, msg)
}

func Warn(fields Fields, msg string) {
	logWithFields("WARN", fields, msg)
}

func logWithFields(level string, fields Fields, msg string) {
	logger := newLogger()
	data := ""
	for k, v := range fields {
		data += fmt.Sprintf("%s=%v;", k, v)
	}
	logger.Println(level, data, msg)
}
