package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Storymate CLI tool", func() {
	var (
		session *gexec.Session
		envs    *envVars
	)

	JustBeforeEach(func() {
		cmd := exec.Command(storymateBinary)
		cmd.Env = envs.toStringArray()

		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {

		envs = &envVars{trackerApiKey: "TEST_API_KEY", trackerProjectID: "TEST_PROJECT_ID"}
	})

	Context("when the app is not configured correctly", func() {
		Context("Tracker API Key", func(){
		BeforeEach(func() {
			envs.trackerApiKey = ""
		})

		It("logs that the API Key is not set", func() {
			Eventually(session.Err).Should(Say("missing TRACKER_API_KEY environment variable"))
		})
		})

		Context("Project ID", func() {
			BeforeEach(func(){
				envs.trackerProjectID = ""
			})

			It("logs that the Project ID is not set", func() {
				Eventually(session.Err).Should(Say("missing TRACKER_PROJECT_ID environment variable"))
			})
		})
	})
})

type envVars struct {
	trackerApiKey      string
	trackerProjectID   string
}

func (e *envVars) toStringArray() []string {
	result := []string{}

	if e.trackerApiKey != "" {
		result = append(result, "TRACKER_API_KEY="+e.trackerApiKey)
	}
	if e.trackerProjectID != "" {
		result = append(result, "TRACKER_PROJECT_ID="+e.trackerProjectID)
	}

	return result
}
