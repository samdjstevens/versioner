package main

import (
	"fmt"
	"os"
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

	// could not get the current version, output the error
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
}

// help command - give some hints on how to use the tool
func help() {
}