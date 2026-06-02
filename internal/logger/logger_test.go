package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := New(INFO)
	if l == nil {
		t.Fatal("New() returned nil")
	}
	if l.level != INFO {
		t.Errorf("New() level = %v, want %v", l.level, INFO)
	}
}

func TestLoggerDebug(t *testing.T) {
	var buf bytes.Buffer
	l := New(DEBUG)
	l.SetOutput(&buf)

	l.Debug("test message")

	output := buf.String()
	if !strings.Contains(output, "DEBUG") {
		t.Errorf("Debug() output should contain 'DEBUG', got %v", output)
	}
	if !strings.Contains(output, "test message") {
		t.Errorf("Debug() output should contain 'test message', got %v", output)
	}
}

func TestLoggerInfo(t *testing.T) {
	var buf bytes.Buffer
	l := New(INFO)
	l.SetOutput(&buf)

	l.Info("info message")

	output := buf.String()
	if !strings.Contains(output, "INFO") {
		t.Errorf("Info() output should contain 'INFO', got %v", output)
	}
	if !strings.Contains(output, "info message") {
		t.Errorf("Info() output should contain 'info message', got %v", output)
	}
}

func TestLoggerWarn(t *testing.T) {
	var buf bytes.Buffer
	l := New(WARN)
	l.SetOutput(&buf)

	l.Warn("warning message")

	output := buf.String()
	if !strings.Contains(output, "WARN") {
		t.Errorf("Warn() output should contain 'WARN', got %v", output)
	}
	if !strings.Contains(output, "warning message") {
		t.Errorf("Warn() output should contain 'warning message', got %v", output)
	}
}

func TestLoggerError(t *testing.T) {
	var buf bytes.Buffer
	l := New(ERROR)
	l.SetOutput(&buf)

	l.Error("error message")

	output := buf.String()
	if !strings.Contains(output, "ERROR") {
		t.Errorf("Error() output should contain 'ERROR', got %v", output)
	}
	if !strings.Contains(output, "error message") {
		t.Errorf("Error() output should contain 'error message', got %v", output)
	}
}

func TestLogLevel(t *testing.T) {
	var buf bytes.Buffer
	l := New(WARN)
	l.SetOutput(&buf)

	// These should not be logged
	l.Debug("debug")
	l.Info("info")

	// These should be logged
	l.Warn("warn")
	l.Error("error")

	output := buf.String()
	if strings.Contains(output, "debug") {
		t.Errorf("Debug message should not be logged at WARN level")
	}
	if strings.Contains(output, "info") {
		t.Errorf("Info message should not be logged at WARN level")
	}
	if !strings.Contains(output, "warn") {
		t.Errorf("Warn message should be logged at WARN level")
	}
	if !strings.Contains(output, "error") {
		t.Errorf("Error message should be logged at WARN level")
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected LogLevel
	}{
		{"DEBUG", DEBUG},
		{"debug", DEBUG},
		{"INFO", INFO},
		{"info", INFO},
		{"WARN", WARN},
		{"warn", WARN},
		{"ERROR", ERROR},
		{"error", ERROR},
		{"invalid", INFO}, // Default to INFO
	}

	for _, tt := range tests {
		result := ParseLogLevel(tt.input)
		if result != tt.expected {
			t.Errorf("ParseLogLevel(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestLoggerSetPrefix(t *testing.T) {
	var buf bytes.Buffer
	l := New(INFO)
	l.SetOutput(&buf)
	l.SetPrefix("TEST")

	l.Info("message")

	output := buf.String()
	if !strings.Contains(output, "TEST") {
		t.Errorf("Output should contain prefix 'TEST', got %v", output)
	}
}
