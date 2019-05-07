package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"os/exec"
)

var _ = Describe("Storymate CLI tool", func() {
	var (
		session            *gexec.Session
		trackerAPIServer   *ghttp.Server
		envs               *envVars
		commandFlag        string
	)

	BeforeEach(func() {
		trackerAPIServer = ghttp.NewServer()
		commandFlag = ""
		envs = &envVars{trackerAPIKey: "TEST_API_KEY", trackerProjectID: "TEST_PROJECT_ID", trackerServerURL: trackerAPIServer.URL()}
	})

	JustBeforeEach(func() {
		cmd := exec.Command(storymateBinary, commandFlag)
		cmd.Env = envs.toStringArray()

		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})


	AfterEach(func() {
		trackerAPIServer.Close()
	})

	Context("help text/usage", func() {
		Context("--help", func() {
			BeforeEach(func() {
				commandFlag = "--help"
			})

			It("displays the help text when the --help flag is present", func() {
				Eventually(session.Err).Should(Say("Usage"))
				Eventually(session.Err).Should(Say("Requires"))
				Consistently(session).ShouldNot(Say("Fetching stories from Pivotal Tracker..."))
			})
		})
		Context("-h", func() {
			BeforeEach(func() {
				commandFlag = "-h"
			})

			It("displays the help text when the -h flag is present", func() {
				Eventually(session.Err).Should(Say("Usage"))
				Eventually(session.Err).Should(Say("Requires"))
				Consistently(session).ShouldNot(Say("Fetching stories from Pivotal Tracker..."))
			})
		})
	})

	Context("when the app is not configured correctly", func() {
		Context("Tracker API Key", func(){
			BeforeEach(func() {
				envs.trackerAPIKey = ""
			})

			It("logs that the API Key is not set", func() {
				Eventually(session.Err).Should(Say("missing TRACKER_API_KEY environment variable"))
				Consistently(session).ShouldNot(Say("Fetching stories from Pivotal Tracker..."))
			})
		})

		Context("Project ID", func() {
			BeforeEach(func(){
				envs.trackerProjectID = ""
			})

			It("logs that the Project ID is not set", func() {
				Eventually(session.Err).Should(Say("missing TRACKER_PROJECT_ID environment variable"))
				Consistently(session).ShouldNot(Say("Fetching stories from Pivotal Tracker..."))
			})
		})
	})

	Context("when the app is configured correctly", func() {

		BeforeEach(func() {
			trackerAPIServer.AppendHandlers(
				ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/services/v5/projects/"+envs.trackerProjectID+"/stories", "with_state=started"),
				ghttp.VerifyHeaderKV("X-TrackerToken", envs.trackerAPIKey),
				ghttp.RespondWith(http.StatusOK, `[{"id": 155484889, "name": "The Beauty and the Beast"},{"id": 155484559, "name": "The Beauty and the Beat"}]`),
				),
			)
		})

		It("display stories IDs that the user owns", func() {
			Eventually(session).Should(Say("Fetching stories from Pivotal Tracker..."))
			Eventually(trackerAPIServer.ReceivedRequests()).Should(HaveLen(1))
			Eventually(session).Should(Say(`1\) 155484889 | The Beauty and the Beast`))
			Eventually(session).Should(Say(`2\) 155484559 | The Beauty and the Beat`))
		})
	})
})

type envVars struct {
	trackerAPIKey      string
	trackerProjectID   string
	trackerServerURL   string
}

func (e *envVars) toStringArray() []string {
	result := []string{}

	if e.trackerAPIKey != "" {
		result = append(result, "TRACKER_API_KEY="+e.trackerAPIKey)
	}
	if e.trackerProjectID != "" {
		result = append(result, "TRACKER_PROJECT_ID="+e.trackerProjectID)
	}
	if e.trackerServerURL != "" {
		result = append(result, "TRACKER_SERVER_URL="+e.trackerServerURL)
	}

	return result
}
