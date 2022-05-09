package cmda

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}

var _ = Describe("test cmda", func() {
	It("should work", func() {
		Expect(A()).To(Succeed())
	})
})
