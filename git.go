package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// use git to get the latest tag for the current directory/project
func getLatestVersionFromGit() (Version, error) {
	var version Version

	// run the git command to get the latest tag
	outputBytes, err := exec.Command("sh", "-c", "git describe --tags --abbrev=0").Output()

	// if the command failed to execute, exit with an error
	if err != nil {
		return version, fmt.Errorf("Failed to get latest version tag from git.")
	}

	// convert to a string and trim new lines
	versionString := strings.Trim(string(outputBytes), "\n")

	// get a Version struct from the tag string
	version, err = NewVersionFromString(versionString)

	// the tag wasn't a valid semver version
	if (err != nil) {
		return version, err
	}

	return version, nil
}