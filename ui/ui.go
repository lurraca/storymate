package ui

import (
	"strings"
	"fmt"
	. "github.com/lurraca/storymate/models"
)


func FormattedStories(stories []Story) string{
	var stringBuilder strings.Builder
	for index, story := range stories {
		stringBuilder.WriteString(fmt.Sprintf("%d) #%d | %s\n", index + 1, story.Id, story.Name))
	}
	return stringBuilder.String()
}
