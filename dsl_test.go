package headandshoulders_test

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/xoebus/headandshoulders"
)

var _ = Describe("DSL", func() {
	Describe("helps you make sure all flakiness is gone", func() {
		RIt("by running repeated tests 3 times by default", func() {
			Expect(rand.Intn(100)).To(BeNumerically(">", 80))
		})

		RIt("but optionally many times for those more tricky ones", func() {
			Expect(rand.Intn(100)).To(BeNumerically(">", 80))
		}, 15)
	})
})

func TestHeadAndShoulders(t *testing.T) {
	// This makes only one of the random numbers be greater than 80.
	rand.Seed(1)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Head and Shoulders Suite")
}
