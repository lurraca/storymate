package ui_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/lurraca/storymate/ui"
  . "github.com/lurraca/storymate/models"
)

var _ = Describe("UI", func() {
	var stories []Story
	Context("when given stories", func(){
		var stories []Story
		BeforeEach(func() {
			stories = []Story{
				Story{Id: 123, Name: "The Beauty and the beast", URL: "https://example.com"},
				Story{Id: 456, Name: "The Beauty and the beat", URL: "https://example.com"},
			}
		})
		It("return a string with formatted stories", func() {
			formattedStories := ui.FormattedStories(stories)
			Expect(formattedStories).To(ContainSubstring("1) #123 | The Beauty and the beast\n2) #456 | The Beauty and the beat"))
		})
	})

	Context("when there are no stories", func() {
		It("returns an empty string", func() {
			formattedStories := ui.FormattedStories(stories)
			Expect(formattedStories).To(Equal(""))
		})
	})
})
