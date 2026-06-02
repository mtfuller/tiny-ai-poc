package version

import "fmt"

var (
	// Version is the current version of the application
	Version = "dev"
	// Commit is the git commit hash
	Commit = "none"
	// BuildDate is the date the binary was built
	BuildDate = "unknown"
)

// GetVersion returns a formatted version string
func GetVersion() string {
	return fmt.Sprintf("%s (commit: %s, built: %s)", Version, Commit, BuildDate)
}

// GetShortVersion returns just the version number
func GetShortVersion() string {
	return Version
}
