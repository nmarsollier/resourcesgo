package logx

import (
	"context"
	"fmt"
	"log"
	"os"
)

// Logger factory
var newLogger = func() *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func ErrorStr(ctx context.Context, err string) {
	logWithFields("ERROR", ctx, err)
}

func Error(ctx context.Context, err error) {
	logWithFields("ERROR", ctx, err.Error())
}

func Info(ctx context.Context, msg string) {
	logWithFields("INFO", ctx, msg)
}

func Warn(ctx context.Context, msg string) {
	logWithFields("WARN", ctx, msg)
}

func logWithFields(level string, ctx context.Context, msg string) {
	logger := newLogger()
	data := ""
	for k, v := range CtxFields(ctx) {
		data += fmt.Sprintf("%s=%v;", k, v)
	}
	logger.Println(level, data, msg)
}
