package main

import (
	"fmt"
	"log"
	"os"
	"flag"
)

func main() {
	validateFlags()
	validateEnvVars()
}

func validateFlags() {
	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.Usage = func() {
        fmt.Fprintf(CommandLine.Output(), usageText())
			}
	flag.Parse()
}

func validateEnvVars() {
	trackerApiKey := os.Getenv("TRACKER_API_KEY")
	trackerProjectID := os.Getenv("TRACKER_PROJECT_ID")

	if trackerApiKey == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_API_KEY environment variable"))
	}

	if trackerProjectID == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_PROJECT_ID environment variable"))
	}
}

func usageText() string {
	var usageText = `storymate is a CLI tool to add the Pivotal Tracker Story ID to the ~/.gitmessage

Usage:
		storymate [--help]

Requires:
		-'TRACKER_API_KEY' environment variable to contain the the user Pivotal Tracker account API Key.
		-'TRACKER_PROJECT_ID' environment variable to contain the Pivotal Tracker project ID that the results will be scoped for.
`

	return usageText
}
