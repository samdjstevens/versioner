package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	// Get the command to run from the first cli arg
	var command string
	args := os.Args[1:]
	if (len(args) == 0) {
		command = ""
	} else {
		command = args[0]
	}

	// Call the appropriate handler function, defaulting to showing the tool info
	switch command {
	case "latest":
		latest()
	case "bump":
		bump()
	case "help":
		help()
	default:
		showInfo()
	}
}

// show the tool info & command list
func showInfo() {
	BlueLn("Versioner")
	fmt.Println()
	YellowLn("Usage:")
	BlueLn("  command [options] [arguments]")
	fmt.Println()
	YellowLn("Available commands:")
	BlueLn("  latest")
	BlueLn("  bump")
	BlueLn("  help")
}

// show the latest/current tag
func latest() {
	// get the latest version
	version, err := getLatestVersionFromGit()

	// could not get the latest version, output the error
	if err != nil {
		RedLn("Error! " + err.Error())
		return
	}

	// print it to the screen
	fmt.Print("The latest version is ")
	GreenLn(version.toString(false))
}

// bump up the latest/current version
func bump() {
	// get the latest version
	version, err := getLatestVersionFromGit()

	// could not get the latest version, output the error
	if err != nil {
		RedLn("Error! " + err.Error())
		return
	}

	// print out the latest version
	YellowLn("Latest version is " + version.toString(false))

	// get which number of the version to bump (major/minor/patch) from the second cli arg
	// default to bumping the patch number
	flag := os.Args[2:]
	var numberToBump string
	if (len(flag) == 0) {
		numberToBump = "patch"
	} else {
		numberToBump = strings.TrimLeft(flag[0], "--")
	}

	// bump up the version returning a new Version struct
	var bumped Version
	switch numberToBump {
	case "major":
		bumped = version.nextMajorVersion()
	case "minor":
		bumped = version.nextMinorVersion()
	case "patch":
		bumped = version.nextPatchVersion()
	default:
		bumped = version.nextPatchVersion()
	}

	// output what the bumped version is
	fmt.Print("Next version will be ")
	GreenLn(bumped.toString(true))

	// ask if they want the git tag created or not
	fmt.Println("Create new git tag with this version? [y/n]")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		err := setVersion(bumped)
		if err != nil {
			RedLn("Error! Failed to add new tag.")
			return
		}
		GreenLn("Git tag created!")
	}
}

// help command - give some hints on how to use the tool
func help() {
}