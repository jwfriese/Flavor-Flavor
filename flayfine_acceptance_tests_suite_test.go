package flayfine_acceptance_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFlayfineAcceptanceTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlayfineAcceptanceTests Suite")
}
