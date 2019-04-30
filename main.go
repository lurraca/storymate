package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	trackerApiKey := os.Getenv("TRACKER_API_KEY")
	if trackerApiKey == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_API_KEY environment variable"))
	}
}
