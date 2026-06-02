package tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()
	cmdArgs := append([]string{"run", "../main.go"}, args...)
	cmd := exec.Command("go", cmdArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	return out.String(), cmd.Run()
}

func TestCLIVersion(t *testing.T) {
	out, err := runCLI(t, "version")
	if err != nil {
		t.Fatalf("version command failed: %v\nOutput: %s", err, out)
	}
	if !strings.Contains(out, "tiny-ai") {
		t.Errorf("expected 'tiny-ai' in output, got: %s", out)
	}
	if !strings.Contains(out, "Version:") {
		t.Errorf("expected 'Version:' in output, got: %s", out)
	}
}

func TestCLIVersionShort(t *testing.T) {
	out, err := runCLI(t, "version", "--short")
	if err != nil {
		t.Fatalf("version --short failed: %v\nOutput: %s", err, out)
	}
	if strings.TrimSpace(out) != "dev" {
		t.Errorf("expected 'dev', got: %s", out)
	}
}

func TestCLIHelp(t *testing.T) {
	out, err := runCLI(t, "--help")
	if err != nil {
		t.Fatalf("--help failed: %v\nOutput: %s", err, out)
	}
	if !strings.Contains(out, "Usage:") {
		t.Errorf("expected 'Usage:' in help output, got: %s", out)
	}
	if !strings.Contains(out, "Available Commands:") {
		t.Errorf("expected 'Available Commands:' in help output, got: %s", out)
	}
}

func TestCLIVerboseFlag(t *testing.T) {
	out, err := runCLI(t, "version", "-v")
	if err != nil {
		t.Fatalf("version -v failed: %v\nOutput: %s", err, out)
	}
	if !strings.Contains(out, "DEBUG") {
		t.Errorf("expected DEBUG log in verbose output, got: %s", out)
	}
}
