package pkgb

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nunnatsa/issue/pkg/pkgd"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}

var _ = Describe("test package B", func() {
	It("should work", func() {
		Expect(B()).To(Succeed())
	})

	It("E should work", func() {
		Expect(pkgd.D()).To(Succeed())
	})
})
