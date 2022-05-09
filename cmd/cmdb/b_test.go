package cmdb

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}

var _ = Describe("test cmdb", func() {
	It("should work", func() {
		Expect(B()).To(Succeed())
	})
})
