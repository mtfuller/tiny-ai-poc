package version

import (
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Save original values
	origVersion := Version
	origCommit := Commit
	origBuildDate := BuildDate

	// Set test values
	Version = "1.0.0"
	Commit = "abc123"
	BuildDate = "2024-01-01"

	defer func() {
		// Restore original values
		Version = origVersion
		Commit = origCommit
		BuildDate = origBuildDate
	}()

	result := GetVersion()
	expected := "1.0.0 (commit: abc123, built: 2024-01-01)"

	if result != expected {
		t.Errorf("GetVersion() = %v, want %v", result, expected)
	}
}

func TestGetShortVersion(t *testing.T) {
	// Save original value
	origVersion := Version

	// Set test value
	Version = "1.0.0"

	defer func() {
		// Restore original value
		Version = origVersion
	}()

	result := GetShortVersion()

	if result != "1.0.0" {
		t.Errorf("GetShortVersion() = %v, want %v", result, "1.0.0")
	}
}

func TestGetVersionContainsComponents(t *testing.T) {
	Version = "1.2.3"
	Commit = "def456"
	BuildDate = "2024-12-19"

	result := GetVersion()

	if !strings.Contains(result, Version) {
		t.Errorf("GetVersion() should contain version %v", Version)
	}
	if !strings.Contains(result, Commit) {
		t.Errorf("GetVersion() should contain commit %v", Commit)
	}
	if !strings.Contains(result, BuildDate) {
		t.Errorf("GetVersion() should contain build date %v", BuildDate)
	}
}
