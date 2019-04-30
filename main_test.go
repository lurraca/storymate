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
	)

	JustBeforeEach(func() {
		cmd := exec.Command(storymateBinary)
		// cmd.Env = envs.toStringArray()

		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})


	Context("when the app is not configured correctly", func() {
		It("logs that the API KEY is not set", func() {
			Eventually(session.Err).Should(Say("missing TRACKER_API_KEY environment variable"))
		})
	})
})

