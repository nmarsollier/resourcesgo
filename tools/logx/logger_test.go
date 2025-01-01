package logx

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	// Save original logger function pointer
	var logger = newLogger
	defer func() { newLogger = logger }()

	// Create a new logger that writes to a buffer
	var buf bytes.Buffer
	var tLogger = newLogger()
	tLogger.SetOutput(&buf)

	// Override original logger function factory
	newLogger = func() *log.Logger {
		return tLogger
	}

	// Execute tests
	f()

	return buf.String()
}

func TestErrorStr(t *testing.T) {

	fields := Fields{"key": "value"}
	output := captureOutput(func() {
		ErrorStr(fields, "test error")
	})

	if !contains(output, "ERROR key=value; test error") {
		t.Errorf("Expected log output to contain 'ERROR key=value; test error', got %s", output)
	}
}

func TestError(t *testing.T) {
	fields := Fields{"key": "value"}
	err := fmt.Errorf("test error")
	output := captureOutput(func() {
		Error(fields, err)
	})
	if !contains(output, "ERROR key=value; test error") {
		t.Errorf("Expected log output to contain 'ERROR key=value; test error', got %s", output)
	}
}

func TestInfo(t *testing.T) {
	fields := Fields{"key": "value"}
	output := captureOutput(func() {
		Info(fields, "test info")
	})
	if !contains(output, "INFO key=value; test info") {
		t.Errorf("Expected log output to contain 'INFO key=value; test info', got %s", output)
	}
}

func TestWarn(t *testing.T) {
	fields := Fields{"key": "value"}
	output := captureOutput(func() {
		Warn(fields, "test warn")
	})
	if !contains(output, "WARN key=value; test warn") {
		t.Errorf("Expected log output to contain 'WARN key=value; test warn', got %s", output)
	}
}

func contains(output, substring string) bool {
	return strings.Contains(output, substring)
}
