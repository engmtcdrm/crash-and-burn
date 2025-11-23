package app

import "regexp"

const (
	// Application Name
	Name = "crash-and-burn"

	// Application Description
	Description = "A simple utility for randomly generating success and failure return codes"
)

var (
	// Application Version
	Version = "dev"
)

// SemVersion returns the semantic version of the application version
// If the version is not in semantic versioning format, it returns the
// original value
func SemVersion() string {
	// Define the regular expression for semantic versioning
	re := regexp.MustCompile(`^v?(\d+\.\d+\.\d+)$`)

	match := re.FindStringSubmatch(Version)

	// If there's a match return the semantic version
	if len(match) > 1 {
		return match[1]
	}

	// If no match, return the original input
	return Version
}
