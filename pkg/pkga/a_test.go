package pkga

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nunnatsa/issue/pkg/pkgd"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}

var _ = Describe("test package A", func() {
	It("should work", func() {
		Expect(A()).To(Succeed())
	})

	It("D should work", func() {
		Expect(pkgd.D()).To(Succeed())
	})
})
