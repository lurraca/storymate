package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestStorymate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Storymate Suite")
}


var storymateBinary string


var _ = SynchronizedBeforeSuite(func() []byte {
	binaryPath, err := gexec.Build("github.com/lurraca/storymate")
	Expect(err).NotTo(HaveOccurred())

	return []byte(binaryPath)
}, func(binaryPath []byte) {
	storymateBinary = string(binaryPath)
})
