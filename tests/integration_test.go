package tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// TestCLIVersion tests the version command
func TestCLIVersion(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "version")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run version command: %v\nOutput: %s", err, out.String())
	}

	output := out.String()
	if !strings.Contains(output, "starterpack-go-cli") {
		t.Errorf("Version output should contain 'starterpack-go-cli', got: %s", output)
	}
	if !strings.Contains(output, "Version:") {
		t.Errorf("Version output should contain 'Version:', got: %s", output)
	}
}

// TestCLIVersionShort tests the version command with --short flag
func TestCLIVersionShort(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "version", "--short")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run version --short command: %v\nOutput: %s", err, out.String())
	}

	output := strings.TrimSpace(out.String())
	// Should only output the version string
	if output != "dev" {
		t.Errorf("Version --short output should be 'dev', got: %s", output)
	}
}

// TestCLIHelp tests the help command
func TestCLIHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "--help")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run help command: %v\nOutput: %s", err, out.String())
	}

	output := out.String()
	if !strings.Contains(output, "Usage:") {
		t.Errorf("Help output should contain 'Usage:', got: %s", output)
	}
	if !strings.Contains(output, "Available Commands:") {
		t.Errorf("Help output should contain 'Available Commands:', got: %s", output)
	}
}

// TestCLIGreet tests the greet command
func TestCLIGreet(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "greet", "World")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run greet command: %v\nOutput: %s", err, out.String())
	}

	output := out.String()
	if !strings.Contains(output, "Hello, World!") {
		t.Errorf("Greet output should contain 'Hello, World!', got: %s", output)
	}
}

// TestCLICalc tests the calc command
func TestCLICalc(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		expected  string
		shouldErr bool
	}{
		{"add", []string{"calc", "5", "3", "--operation", "add"}, "8", false},
		{"subtract", []string{"calc", "10", "4", "--operation", "subtract"}, "6", false},
		{"multiply", []string{"calc", "6", "7", "--operation", "multiply"}, "42", false},
		{"divide", []string{"calc", "20", "4", "--operation", "divide"}, "5", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "../main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			if tt.shouldErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Failed to run calc command: %v\nOutput: %s", err, out.String())
			}

			output := out.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Calc output should contain '%s', got: %s", tt.expected, output)
			}
		})
	}
}

// TestCLIProcess tests the process command
func TestCLIProcess(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "process")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run process command: %v\nOutput: %s", err, out.String())
	}

	output := out.String()
	if !strings.Contains(output, "Successfully processed") {
		t.Errorf("Process output should contain 'Successfully processed', got: %s", output)
	}
}

// TestCLIVerboseFlag tests the verbose flag
func TestCLIVerboseFlag(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "greet", "Test", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run command with verbose flag: %v\nOutput: %s", err, out.String())
	}

	output := out.String()
	// When verbose is on, debug logs should appear
	if !strings.Contains(output, "DEBUG") {
		t.Errorf("Verbose output should contain 'DEBUG', got: %s", output)
	}
}
