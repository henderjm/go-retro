package feedback_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFeedback(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feedback Suite")
}
