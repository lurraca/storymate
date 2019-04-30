package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	trackerApiKey := os.Getenv("TRACKER_API_KEY")
	trackerProjectID := os.Getenv("TRACKER_PROJECT_ID")
	if trackerApiKey == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_API_KEY environment variable"))
	}
	if trackerProjectID == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_PROJECT_ID environment variable"))
	}
}
