package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"

	"github.com/lurraca/storymate/ui"
	. "github.com/lurraca/storymate/models"
)
var trackerServerURL string
var trackerAPIKey string
var trackerProjectID string

func main() {
	validateFlags()
	validateEnvVars()

	stories := fetchStartedStories()
	fmt.Println(ui.FormattedStories(stories))
}

func validateFlags() {
	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.Usage = func() {
        fmt.Fprintf(CommandLine.Output(), usageText())
			}
	flag.Parse()
}

func validateEnvVars() {
	trackerAPIKey = os.Getenv("TRACKER_API_KEY")
	trackerProjectID = os.Getenv("TRACKER_PROJECT_ID")
	trackerServerURL = os.Getenv("TRACKER_SERVER_URL")

	if trackerAPIKey == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_API_KEY environment variable"))
	}

	if trackerProjectID == "" {
		log.Fatal(fmt.Sprintf("missing TRACKER_PROJECT_ID environment variable"))
	}
	if trackerServerURL == "" {
		trackerServerURL = "https://www.pivotaltracker.com"
	}
}

func fetchStartedStories() []Story {
	fmt.Println("Fetching stories from Pivotal Tracker...")

	req, err := http.NewRequest("GET", trackerServerURL+"/services/v5/projects/"+trackerProjectID+"/stories", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("X-TrackerToken", trackerAPIKey)
	q := req.URL.Query()
	q.Add("with_state", "started")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	var stories []Story

	defer res.Body.Close()
	trackerStoriesJson, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(trackerStoriesJson, &stories)

	if err != nil {
		fmt.Println("There was an error unmarshalling the stories:", err)
		fmt.Println("Body: ", trackerStoriesJson)
	}

	return stories
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
