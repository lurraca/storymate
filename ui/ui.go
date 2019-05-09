package ui

import (
	"strings"
	"fmt"
	. "github.com/lurraca/storymate/models"
)


func PrintFormattedStories(stories []Story) string{
	var stringBuilder strings.Builder
	for index, story := range stories {
		stringBuilder.WriteString(fmt.Sprintf("%d) #%d | %s\n", index + 1, story.Id, story.Name))
	}
	return stringBuilder.String()
}

func FormattedStories(stories []Story) map[int]Story{
	mapOfStories := map[int]Story{}
	for index, story := range stories {
		mapOfStories[index + 1] = story
	}
	return mapOfStories
}
