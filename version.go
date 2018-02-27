package main

import (
	"strconv"
	"strings"
	"fmt"
	"regexp"
)

// struct representing a semver version
type Version struct {
	major int
	minor int
	patch int
}

// return a string representation of the Version, e.g. v8.5.92
func (v *Version) toString(withV bool) string {
	// Form the "x.y.z" string from the major/minor/patch numbers
	parts := [3]string{strconv.Itoa(v.major), strconv.Itoa(v.minor), strconv.Itoa(v.patch)}
	versionString := strings.Join(parts[:], ".")

	// If they want the "v" prefix, then add it
	if (withV) {
		return "v" + versionString
	}

	return versionString
}

// return a new Version with the patch number bumped
func (v *Version) nextPatchVersion() Version {
	return Version{major: v.major, minor: v.minor, patch: v.patch + 1}
}

// return a new Version with the minor number bumped, patch reset to 0
func (v *Version) nextMinorVersion() Version {
	return Version{major: v.major, minor: v.minor + 1, patch: 0}
}

// return a new Version with the major number bumped, minor and patch reset to 0
func (v *Version) nextMajorVersion() Version {
	return Version{major: v.major + 1, minor: 0, patch: 0}
}

// create a new Version struct from a semver version string.
// returns an error if not in x.y.z format
func NewVersionFromString(versionString string) (Version, error) {
	var version Version

	// Check that the version is in the semver format of x.y.z (with optional v prefix)
	regexp := regexp.MustCompile("^v?([0-9]+)\\.([0-9]+)\\.([0-9]+)$")
	matches := regexp.FindStringSubmatch(versionString)

	// If it's not in the x.y.z format, exit with an error
	if (len(matches) != 4) {
		return version, fmt.Errorf("Version string \"%s\" does not seem to follow Semver.", versionString)
	}

	// Convert the string-ints into actual ints
	majorInt, _ := strconv.Atoi(matches[1])
	minorInt, _ := strconv.Atoi(matches[2])
	patchInt, _ := strconv.Atoi(matches[3])

	// Set the new version
	version = Version{major: majorInt, minor: minorInt, patch: patchInt}

	return version, nil
}